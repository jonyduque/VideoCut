# Mapeamento do Código-Fonte (Codebase Map)

Abaixo está a representação da estrutura de pastas e arquivos do projeto Wails, seguida pela explicação das funções de cada um dos componentes.

## Estrutura de Diretórios
```text
fatiador-de-videos/
|-- wails.json
|-- main.go
|-- app.go
|-- frontend/
    |-- index.html
    |-- src/
    |   |-- main.ts
    |   |-- style.css
    |-- wailsjs/
        |-- go/
        |   |-- main/
        |       |-- App.js
        |       |-- App.d.ts
```
		
## Descrição dos Arquivos e Pastas

### Diretório Raiz

O diretório raiz do projeto hospeda os arquivos de configuração do framework e a camada de backend construída em linguagem Go.
1. O arquivo wails.json é responsável por armazenar as configurações fundamentais do projeto Wails, abrangendo dados como o nome do aplicativo, a versão atual e as diretivas de compilação (build).
2. O arquivo main.go atua como o ponto de partida do software. Sua função exclusiva é inicializar a janela nativa do sistema operacional, configurar as opções do ciclo de vida da aplicação e registrar o núcleo lógico do sistema para uso posterior.
3. O arquivo app.go armazena a estrutura central do aplicativo e todos os métodos de backend que precisam ser exportados. É neste local que se encontra o código para capturar os caminhos dos arquivos, invocar o ffprobe para ler metadados e acionar o ffmpeg para executar as divisões da mídia.

### Diretório frontend

Esta pasta e suas subpastas contêm a totalidade da interface gráfica do aplicativo, que é renderizada como uma página web interativa.
1. O arquivo index.html é o esqueleto da interface. Ele desenha os painéis principais do layout e realiza a importação de todas as folhas de estilo e dos arquivos de script necessários.
2. A pasta src é o local de desenvolvimento ativo do frontend. O arquivo main.ts captura as ações do usuário, gerencia os eventos de arrastar e soltar (drag and drop) e dispara os comandos para o backend. O arquivo style.css dita todas as regras visuais, como cores, alinhamentos e respostas visuais ao usuário.
3. A pasta wailsjs é uma estrutura gerada e gerenciada de forma totalmente automática pelo framework Wails. Ela fornece as definições de tipo do TypeScript e as pontes de comunicação em JavaScript necessárias para que os botões do frontend consigam acionar os métodos escritos no arquivo app.go.