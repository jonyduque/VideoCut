# Especificações de Design e Interface

A interface deve ser minimalista, focada na usabilidade direta, funcionando em uma janela de tamanho fixo para evitar complexidade de redimensionamento de layout.

**Área de Recepção de Arquivo (Drag and Drop)**
O topo da interface será dominado por uma área delimitada por bordas tracejadas. O texto indicará a ação de arrastar o vídeo para este local ou clicar para abrir o seletor nativo do Windows. Durante o evento de "dragover", a área mudará de cor para fornecer feedback visual ao usuário.

**Painel de Metadados e Corte**
Abaixo da área de recepção, após o carregamento do arquivo, o nome do vídeo e sua duração total serão exibidos. Um controle deslizante (slider) horizontal ocupará a largura da janela, permitindo deslizar do segundo zero até o final do vídeo. O tempo exato selecionado no slider será exibido em um texto atualizado em tempo real.

**Configurações de Saída**
Uma seção dedicada conterá os campos de entrada de texto para as configurações finais.
1. Um campo numérico para a "Sobra" (em segundos).
2. Dois campos de texto para os nomes do "Vídeo Parte 1" e "Vídeo Parte 2".
3. Uma caixa de seleção (checkbox) com o texto "Salvar na mesma pasta do arquivo original". Se desmarcada, exibirá um botão ao lado para escolher o diretório de destino.

**Controle de Ação**
Na parte inferior, um botão principal de destaque será responsável por iniciar o corte. Durante o processamento, este botão será desabilitado e seu texto mudará para indicar o progresso, evitando execuções duplicadas.