# Stack Tecnológica: VideoCut (Fatiador de Vídeos)

## 1. Backend (Core)
- **Linguagem:** Go (Golang) para lógica de sistema e I/O.
- **Framework:** **Wails (v2/v3)** para unificação de backend nativo e interface web.
- **Processamento de Mídia:** **FFmpeg** e **ffprobe** (Binários estáticos compilados para Windows).
- **Integração com SO:** Pacote `os/exec` para chamadas de sistema e `path/filepath` para gestão de caminhos relativos.

## 2. Frontend (Interface)
- **Linguagem:** TypeScript para tipagem forte e segurança de dados.
- **Framework:** **React** para construção de componentes modulares.
- **Estilização:** **Vanilla CSS** puro para máximo controle visual e performance, seguindo as diretrizes de design minimalista.
- **Gestão de Estado:** **Zustand** para um gerenciamento de estado leve, rápido e escalável.

## 3. Comunicação (IPC)
- **Wails Bindings:** Geração automática de pontes TypeScript para chamar funções Go.
- **Eventos:** Uso de eventos do sistema Wails para feedback em tempo real (ex: progresso de processamento).

## 4. Ferramentas de Desenvolvimento
- **Gerenciador de Pacotes:** `npm` ou `pnpm` para dependências do frontend.
- **Compilação:** `wails build` para gerar o executável final auto-contido.
- **Controle de Versão:** Git.

## 5. Infraestrutura e Deployment
- **Portabilidade:** Executável único (.exe) sem dependências externas de runtime (além do WebView2 nativo do Windows).
- **Distribuição:** Pacote ZIP contendo o executável e a pasta de binários do FFmpeg.