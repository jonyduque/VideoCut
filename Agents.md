# Instruções para Agentes de IA (SPD)

Este documento define os papéis, o contexto operacional para as interações com Modelos de Linguagem durante a fase de codificação do projeto e serve como um índice do escopo arquitetural.

## Contexto Geral e Arquitetura para todos os Agentes
O sistema é um utilitário desktop para Windows construído sob o modelo de aplicação de página única (SPA) com um backend local nativo, unificado pelo framework Wails. A ferramenta divide vídeos utilizando os binários do FFmpeg e ffprobe, que devem residir na mesma pasta do executável final. O software tem acoplamento zero com o sistema operacional: não exige instalação, não utiliza o registro do Windows, não requer elevação de privilégios (UAC) e não depende de bibliotecas externas complexas. A comunicação entre as camadas ocorre via Inter-Process Communication (IPC) providenciada pelo Wails.

## Resumo das Documentações do Projeto
Para garantir a coesão do código e o alinhamento com a arquitetura definida, o agente de IA deve basear suas respostas no conteúdo dos seguintes documentos de referência:

O documento Codebase_map.md contém a árvore literal de diretórios do projeto e a explicação detalhada do propósito de cada arquivo. Ele deve ser consultado sempre que houver necessidade de entender onde posicionar uma nova lógica ou como as pastas do frontend e backend interagem.

O documento Design.md especifica as diretrizes da interface gráfica. Ele descreve a disposição dos elementos na tela, o comportamento visual esperado para a área de drag and drop, a exibição de metadados e as mudanças de estado dos botões durante o processamento.

O documento Project.md detalha a arquitetura do sistema, o fluxo de dados (IPC) entre Go e TypeScript e as fases exatas de implementação do software. Ele orienta o trabalho desde o setup inicial, passando pela construção da lógica I/O em Go, integração com o frontend, até a etapa final de testes de resiliência.

O documento Readme.md funciona como o manual do usuário final e guia de compilação. Ele contém as instruções de uso prático, os pré-requisitos absolutos do sistema e os comandos necessários para compilar o executável no ambiente de desenvolvimento.

## Agente de Backend (Especialista em Go)
Sua função é escrever o núcleo lógico de I/O no arquivo app.go. Você deve focar estritamente no uso do pacote os/exec para chamadas seguras e em segundo plano ao FFmpeg e ao ffprobe. Utilize sempre os pacotes path/filepath e os.Executable para descobrir o diretório de execução em tempo real, garantindo que a resolução dos binários seja sempre relativa ao nosso executável. A lógica de manipulação de arquivos não deve falhar caso o diretório de destino selecionado seja o mesmo do arquivo original. Você será responsável por expor as funções que retornarão as promessas (Promises) aguardadas pelo frontend.

## Agente de Frontend (Especialista em TypeScript e DOM)
Sua função é escrever a camada de apresentação utilizando HTML e CSS nativos, atuando como o controlador de estado da interface em TypeScript. Você deve implementar o listener de drag and drop interceptando os eventos de tela padrão do navegador. Toda a comunicação com o ambiente Go deve ser executada exclusivamente consumindo as funções estáticas geradas na pasta wailsjs. A interface deverá obrigatoriamente bloquear os inputs do usuário e alterar o estado visual para "carregando" enquanto aguarda a resolução da Promise de processamento (IPC) enviada pelo backend.

# PRD: VideoCut

**1. Visão Geral do Produto**
O aplicativo é uma ferramenta de desktop local para o sistema operacional Windows, desenhada especificamente para particionar arquivos de vídeo. O foco absoluto é a simplicidade de operação e a portabilidade. O software não possuirá instalador e funcionará de forma independente a partir de qualquer pasta onde for alocado, desde que o executável original do FFmpeg esteja localizado no mesmo diretório. 

**2. Arquitetura e Stack Tecnológico**
O sistema seguirá o modelo cliente-servidor embutido.
O núcleo da aplicação será escrito em Go. Ele será responsável por lidar com o sistema de arquivos, descobrir o caminho do executável em tempo de execução, invocar processos filhos (FFmpeg/ffprobe) e capturar seus retornos.
A interface visual (frontend) será construída via Wails, utilizando HTML, CSS e TypeScript/JavaScript. O frontend enviará os parâmetros inseridos pelo usuário para as funções exportadas do Go de maneira assíncrona, evitando o congelamento da tela durante o processamento do vídeo.

**3. Requisitos Funcionais**
O programa deve garantir a execução das seguintes funcionalidades principais, em ordem de fluxo:

1. Importação de Mídia: Um botão na interface que abre a caixa de diálogo nativa do Windows para a escolha do arquivo de vídeo.
2. Leitura de Metadados: Uma função oculta que executa o comando `ffprobe` no arquivo recém-selecionado para extrair a duração total em segundos.
3. Definição do Corte: Um controle deslizante horizontal (slider) cujo valor máximo é a duração total do vídeo, permitindo ao usuário escolher visualmente o instante da divisão.
4. Ajuste de Sobra: Um campo numérico para inserir o tempo de segurança (em segundos) que servirá de intersecção entre as duas novas mídias geradas.
5. Nomenclatura Personalizada: Dois campos de entrada de texto para definir os nomes de saída dos arquivos resultantes.
6. Processamento: Um botão principal de ação que aciona a lógica de fatiamento.

**4. Requisitos Não Funcionais**
Para manter a filosofia de uma ferramenta utilitária leve, as seguintes premissas técnicas devem ser obedecidas:

1. Zero Instalação: Compilação de um executável único. Nenhuma dependência externa ou framework adicional deve ser exigido na máquina alvo.
2. Contexto Local: A resolução do caminho do FFmpeg deve ser relativa ao executável pai. O código em Go utilizará `os.Executable()` e `filepath.Dir()` para garantir que ele procure o binário na mesma pasta, independentemente de onde o usuário a coloque.
3. Privilégios Básicos: Toda operação de leitura e escrita será contida em pastas de domínio do usuário. Elevação de privilégios via UAC está estritamente proibida.
4. UX Resiliente: A interface deve exibir estados de carregamento (loading) durante o trabalho do FFmpeg e desabilitar botões para prevenir cliques duplos.

**5. Lógica da Sobra (Overlap) e Chamada ao Sistema**
A matemática da sobra garante que nenhuma informação se perca na transição, gerando uma redundância de segurança. Se o usuário definir o ponto de corte no tempo $T_c$ e uma sobra designada como $S$:

O primeiro comando ordenará que o FFmpeg gere o primeiro vídeo do instante $0$ até o momento exato de $T_c + S$.

O segundo comando ordenará que o FFmpeg inicie a captura do segundo vídeo no momento exato de $T_c - S$ e vá até o término da mídia.

Devido à necessidade de garantir a precisão exata das variáveis $T_c$ e $S$, a chamada via pacote `os/exec` no Go deverá, por padrão, omitir o parâmetro de cópia rápida e forçar a re-codificação, evitando os erros de proximidade dos quadros-chave (keyframes). O comando base a ser montado no Go será semelhante a `ffmpeg -i "entrada.mp4" -ss [INICIO] -to [FIM] "saida.mp4"`.

**6. Ponto de Partida para a Implementação**
Para iniciar via método de prompt estruturado, o próximo passo lógico é definir os "Bindings" do Wails. Você deve instruir a criação de um arquivo `app.go` contendo uma *struct* principal. Essa *struct* deve expor métodos como `SelectVideo()`, `GetDuration()` e `SplitVideo()`. Assim que esse contrato de dados estiver pronto, o código TypeScript do frontend já terá os canais exatos para enviar as variáveis do slider e receber as confirmações de conclusão.

**7. Interatividade via Drag and Drop**
A interface principal agora contará com uma zona de recepção de arquivos. O usuário poderá arrastar um vídeo diretamente do Explorador de Arquivos para dentro da janela do programa. Isso exige que o frontend em TypeScript implemente listeners para os eventos de "dragover" e "drop". Ao detectar o arquivo, o caminho absoluto deve ser enviado imediatamente para o backend em Go, que validará a extensão do arquivo e acionará o ffprobe para carregar a duração e atualizar o slider de corte automaticamente.

**8. Flexibilidade de Destino de Saída**
O programa oferecerá um seletor de modo de salvamento com duas opções lógicas. A primeira opção manterá os arquivos resultantes na mesma pasta do arquivo original, o que simplifica a organização para edições rápidas. A segunda opção permitirá que o usuário defina uma pasta específica através de um botão de navegação. Caso a pasta de destino não seja definida, o sistema deve assumir como padrão o diretório onde o executável do programa está rodando para evitar falhas de gravação.

**9. Arquitetura de Gerenciamento de Caminhos**
No backend em Go, o gerenciamento de caminhos será centralizado em uma estrutura que armazena tanto o local do vídeo de entrada quanto o diretório de destino escolhido. Utilizando o pacote path/filepath, o sistema extrairá o diretório pai do arquivo original sempre que um novo vídeo for carregado. Se a opção de "mesma pasta" estiver ativa, as variáveis de saída das chamadas do FFmpeg serão construídas concatenando esse diretório pai com os nomes definidos pelo usuário. Se um diretório customizado for escolhido via runtime.OpenDirectoryDialog do Wails, esse novo caminho passará a ser a base para todas as operações de exportação.

**10. Fluxo de Trabalho e Chamadas de Sistema**
A execução do corte permanecerá resiliente e sem necessidade de elevação de privilégios. Ao disparar o processo, o Go verificará a existência do diretório de destino e, se necessário, tentará criá-lo (caso o usuário tenha permissão na pasta pai). O comando enviado ao FFmpeg será dinamicamente montado para refletir o caminho completo de saída, garantindo que os arquivos finais sejam gravados exatamente onde o usuário espera, sem gerar arquivos temporários em locais protegidos do Windows.

**11. Sugestões para a Implementação Técnica**
Para o recurso de arrastar e soltar, você pode aproveitar que o Wails permite capturar eventos do sistema. No seu arquivo principal de frontend, você deve impedir o comportamento padrão do navegador para o evento de drop e capturar o ```e.dataTransfer.files[0].path```. Como o Wails tem acesso direto ao sistema de arquivos, esse caminho pode ser passado para uma função Go que faz o "bind" com o restante da lógica de metadados.

Para a escolha da pasta, recomendo criar uma variável booleana no estado da aplicação chamada ```useSourceFolder```. Se ela for verdadeira, o backend ignora qualquer caminho de saída salvo e executa ```filepath.Dir(inputFilePath)``` no momento da renderização do vídeo. Isso mantém o código limpo e evita que o usuário precise configurar o destino repetidamente se estiver trabalhando com vários vídeos na mesma pasta.