# VideoCut ✂️

VideoCut é uma ferramenta desktop ultra-portátil para Windows projetada para particionar arquivos de vídeo com precisão e facilidade. Construída com **Go** e **Wails**, ela oferece uma experiência self-contained (auto-contida), eliminando a necessidade de instalar dependências externas como o FFmpeg manualmente.

## ✨ Diferenciais

- **Zero Instalação:** Um único executável. Sem instaladores, sem chaves de registro, sem complicação.
- **FFmpeg Embutido:** Os binários do FFmpeg e FFprobe estão embutidos no software e são extraídos automaticamente para o diretório de execução se não forem encontrados no sistema.
- **Privilégios Mínimos:** Não requer elevação de privilégios (UAC). Funciona inteiramente no espaço do usuário.
- **Lógica de Overlap (Sobra):** Permite definir uma "sobra" de tempo que cria uma intersecção de segurança entre as duas partes geradas, garantindo que nenhum contexto da transição seja perdido.

## 🚀 Como Usar

1. **Importação:** Arraste um vídeo diretamente para a janela ou clique na área de upload para selecionar um arquivo.
2. **Ponto de Corte:** Utilize o slider intuitivo para definir o momento exato da divisão. O tempo é atualizado em tempo real.
3. **Configuração de Sobra:** Ajuste o campo "Sobra" para definir quantos segundos de intersecção as partes terão (ex: 2 segundos de sobra garantem que o final da parte 1 e o início da parte 2 compartilhem o mesmo trecho).
4. **Nomenclatura:** Defina os nomes dos arquivos de saída nos campos correspondentes.
5. **Destino:** Escolha entre salvar na mesma pasta do arquivo original ou selecionar um diretório customizado.
6. **Processar:** Clique em "Cortar Vídeo". A interface exibirá um estado de carregamento enquanto o processamento ocorre em segundo plano.

## 🏗️ Arquitetura Técnica

O VideoCut utiliza o modelo **SPA (Single Page Application)** com um backend nativo em Go:

- **Backend (Go):** Gerencia operações de I/O, descoberta de diretórios em tempo de execução e invocação segura dos processos filhos (FFmpeg/FFprobe) via pacotes `os/exec`.
- **Frontend (TypeScript/React):** Interface visual moderna rodando sobre WebView2, comunicando-se com o backend através de **IPC (Inter-Process Communication)** providenciado pelo framework Wails.
- **Processamento:** Para garantir precisão absoluta, o software força a re-codificação (`libx264`), evitando erros de proximidade de keyframes comuns em cópias rápidas.

## 🛠️ Requisitos de Compilação

Caso deseje compilar o projeto a partir do código-fonte:

- **Go** 1.23+
- **Node.js** e **NPM**
- **Wails CLI** v2
- **WebView2 Runtime** (padrão no Windows 10/11)

**Comando de Build:**
```powershell
wails build -platform windows/amd64 -ldflags="-s -w"
```

## 📄 Licença

Este projeto é uma ferramenta utilitária distribuída sob as diretrizes de portabilidade e baixo acoplamento. Consulte o arquivo `Project.md` para detalhes arquiteturais profundos.

---
Desenvolvido por [jonyduque](https://github.com/jonyduque)
