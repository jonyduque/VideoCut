# Initial Concept

## Descrição Detalhada
O Fatiador de Vídeos é uma ferramenta utilitária de desktop local para o sistema operacional Windows, criada para resolver a necessidade de cortes precisos de mídia de forma rápida e portátil. O diferencial primário do projeto é sua total independência estrutural e baixo acoplamento ao sistema operacional do usuário. 

O software não requer processo de instalação, não escreve chaves no registro do sistema, não exige elevação de privilégios de administrador (UAC) e não possui dependências externas além da presença física dos binários do FFmpeg (ffmpeg.exe e ffprobe.exe) em seu próprio diretório. Ele foi desenhado para ser alocado em qualquer pasta do computador ou dispositivo de armazenamento removível e simplesmente funcionar. O fluxo de trabalho principal atende à necessidade de dividir um arquivo de vídeo em duas partes contínuas, permitindo a configuração de uma "sobra" de tempo (overlap). Esta sobra cria uma área de intersecção segura entre os dois novos arquivos gerados, garantindo que o contexto da transição não seja perdido.

## Arquitetura do Sistema
A arquitetura adota o modelo de aplicação de página única (SPA) operando sobre um backend local nativo, sendo unificada pelo framework Wails. A separação de responsabilidades entre as camadas é rígida, focada em manter a interface fluida enquanto o processamento pesado ocorre no sistema operacional.

**Camada de Backend (Go):**
O núcleo lógico, de segurança e de integração com o sistema operacional é escrito em Go. Esta camada é estritamente encarregada de manipular as operações de entrada e saída de arquivos (I/O). Ela utiliza as bibliotecas padrão da linguagem para descobrir dinamicamente o caminho absoluto de onde a aplicação está rodando, o que garante a localização dos binários do FFmpeg independentemente de onde o usuário salve o programa. Além disso, gerencia a invocação de processos filhos de forma oculta através do pacote `os/exec`. O backend expõe uma interface de comunicação restrita (bindings) que recebe os dados do usuário, executa a lógica matemática de sobreposição temporal, aciona os comandos do sistema e devolve respostas assíncronas de sucesso ou erro.

**Camada de Frontend (TypeScript/HTML/CSS):**
A camada de apresentação é visual e interativa, rodando no motor de renderização leve providenciado pelo Wails (WebView2 no Windows). O TypeScript atua como o controlador de estado da interface, monitorando ativamente os eventos do Document Object Model (DOM), como o evento de arrastar um arquivo de vídeo para a zona designada na tela. Esta camada é responsável por validar previamente os dados inseridos e se comunicar com o backend consumindo as promessas (Promises) geradas pelas funções em Go.

**Fluxo de Dados e Comunicação (IPC):**
A interação entre a interface web e o binário Go ocorre via Inter-Process Communication (IPC) otimizada pelo Wails. Quando o usuário clica em "Cortar", o TypeScript empacota os valores do formulário e aciona a função exportada correspondente. A rotina em Go inicia o processo do FFmpeg e bloqueia o seu retorno até que a re-codificação termine. Durante este período, o TypeScript altera o estado da interface visual para exibir uma tela de carregamento, aguardando a resolução da Promise enviada pelo backend para notificar o usuário e liberar a tela novamente.

## Fases de Implementação
1. Fase 1 - Setup e Preparação Estrutural: Inicialização do projeto base com Wails e o template de TypeScript. Limpeza dos arquivos genéricos de exemplo e construção do esqueleto primário do arquivo app.go para abrigar os métodos exportados.
2. Fase 2 - Desenvolvimento do Backend (Go): Criação das rotinas de resolução de diretórios absolutos e relativos. Implementação da chamada silenciosa ao ffprobe para extração da duração da mídia. Desenvolvimento da lógica matemática de cálculo do ponto de corte considerando o tempo de "sobra" e montagem do comando complexo do ffmpeg forçando a recodificação frame a frame.
3. Fase 3 - Desenvolvimento do Frontend (Interface): Construção da marcação semântica em HTML seguindo as especificações do Design.md. Implementação das regras visuais em CSS e da lógica em TypeScript para interceptação nativa do evento de drag and drop. Amarração dos inputs e botões com as pontes de comunicação em wailsjs.
4. Fase 4 - Testes e Validação de Resiliência: Execução de baterias de teste simulando a ausência completa de privilégios administrativos. Validação rigorosa do comportamento do sistema quando a pasta de destino selecionada for idêntica ao diretório de origem. Verificação final da precisão dos cortes gerados aplicando variações múltiplas no tempo de intersecção.

---

# Product Definition: VideoCut (Fatiador de Vídeos)

## 1. Product Vision
VideoCut is a portable, zero-installation desktop utility for Windows designed for fast and precise video splitting. Its core differentiator is total independence: it runs without an installer, registry modifications, or administrative privileges (UAC), relying on bundled FFmpeg binaries located in its execution directory.

## 2. Target Audience
- **Content Creators:** Users needing to quickly segment long recordings for social media or archives.
- **Privacy-Conscious Users:** Those who prefer local processing over cloud-based tools.
- **Utility Seekers:** Users requiring a "run anywhere" tool that can be kept on a USB drive.

## 3. Core Value Proposition
- **Portability:** Single executable that works anywhere.
- **Context Preservation:** The "Overlap/Sobra" feature ensures transitions aren't lost between segments.
- **Speed & Precision:** Direct FFmpeg invocation with frame-accurate re-encoding.
- **Simplicity:** Drag-and-drop workflow with intuitive visual feedback.

## 4. Key Functional Requirements
- **Import:** Supports drag-and-drop and native file selection for universal video formats.
- **Metadata Extraction:** Automatic duration retrieval via `ffprobe`.
- **Visual Cutting:** Timeline slider for selecting the split point ($T_c$).
- **Overlap (Sobra):** Configurable intersection time ($S$) to create redundancy between Part 1 ($0$ to $T_c + S$) and Part 2 ($T_c - S$ to end).
- **Architecture for Growth:** Designed to support multi-part splitting in future versions.
- **Flexible Output:** Option to save in the source folder or a custom destination.

## 5. User Experience (UX) Principles
- **Minimalist Design:** Fixed-window, distraction-free interface.
- **Active Feedback:** Visual cues for drag-and-drop, real-time slider updates, and loading states during processing.
- **Error Resilience:** Graceful handling of missing binaries or restricted folder permissions.

## 6. Business & Distribution Strategy
- **Self-Contained Distribution:** Bundles `ffmpeg.exe` and `ffprobe.exe` within the application folder.
- **No Installation Required:** Optimized for "download and run" scenarios.
- **Open Architecture:** Built using Wails (Go + Modern Web Framework) for future cross-platform potential (initially targeted at Windows).
