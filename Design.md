# Especifica??es de Design e Interface

A interface deve ser minimalista, focada na usabilidade direta, funcionando em uma janela de tamanho fixo para evitar complexidade de redimensionamento de layout.

**?rea de Recep??o de Arquivo (Drag and Drop)**
O topo da interface ser? dominado por uma ?rea delimitada por bordas tracejadas. O texto indicar? a a??o de arrastar o v?deo para este local ou clicar para abrir o seletor nativo do Windows. Durante o evento de "dragover", a ?rea mudar? de cor para fornecer feedback visual ao usu?rio.

**Painel de Metadados e Corte**
Abaixo da ?rea de recep??o, ap?s o carregamento do arquivo, o nome do v?deo e sua dura??o total ser?o exibidos. Um controle deslizante (slider) horizontal ocupar? a largura da janela, permitindo deslizar do segundo zero at? o final do v?deo. O tempo exato selecionado no slider ser? exibido em um texto atualizado em tempo real.

**Configura??es de Sa?da**
Uma se??o dedicada conter? os campos de entrada de texto para as configura??es finais.
1. Um campo num?rico para a "Sobra" (em segundos).
2. Dois campos de texto para os nomes do "V?deo Parte 1" e "V?deo Parte 2".
3. Uma caixa de sele??o (checkbox) com o texto "Salvar na mesma pasta do arquivo original". Se desmarcada, exibir? um bot?o ao lado para escolher o diret?rio de destino.

**Controle de A??o**
Na parte inferior, um bot?o principal de destaque ser? respons?vel por iniciar o corte. Durante o processamento, este bot?o ser? desabilitado e seu texto mudar? para indicar o progresso, evitando execu??es duplicadas.
