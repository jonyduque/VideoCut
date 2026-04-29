# Fatiador de Vídeos

Um utilitário de desktop leve e portátil para particionar arquivos de vídeo com precisão, construído com Go e Wails.

**Sobre o Projeto**
Este aplicativo foi desenvolvido para simplificar o processo de dividir um vídeo em duas partes, permitindo configurar um tempo de "sobra" (overlap) entre os dois novos arquivos. Todo o processamento é feito localmente, sem necessidade de internet, instalação ou privilégios de administrador.

**Pré-requisitos de Execução**
Para que o programa funcione corretamente, ele depende dos binários do FFmpeg.
1. O arquivo executável principal (fatiador.exe).
2. O arquivo ffmpeg.exe na mesma pasta do programa.
3. O arquivo ffprobe.exe na mesma pasta do programa.

**Como Usar**
1. Abra o executável do programa.
2. Arraste um arquivo de vídeo para dentro da interface ou clique na área tracejada para selecionar.
3. Use o controle deslizante para escolher o ponto exato da divisão.
4. Defina os segundos de "sobra" para garantir que a transição não perca informações importantes.
5. Escolha os nomes dos novos arquivos e o diretório de destino.
6. Clique em cortar e aguarde a finalização do processo.

**Compilação (Para Desenvolvedores)**
Certifique-se de ter o Go e o Wails instalados na sua máquina. Navegue até a raiz do projeto e execute o comando de build do Wails para gerar o binário único do Windows.