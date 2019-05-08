# Introdução
Repositório para o desenvolvimento do projeto da disciplina de sistemas de programação

# Processador hospedeiro
O processador hospedeiro é uma máquina virtual sobre a qual será executado o simulador.
Sua vantagem é a simplicidade de construção e sua desvantagem é a perda de eficiência pelo simulador estar rodando sobre um software.
As características de memória, registradores<!--, equipamentos de entrada, saída e armazenamento,--> e os blocos funcionais estão descritos nos próximos itens.

## Memória
<!-- Já pode ser implentado -->
* 16 bits por palavra
* 16 bits (2 bytes) de endereçamento
* número máximo de palavras: 2<sup>16</sup> = 65536 bytes = 32768 palavras

## Registradores
<!-- Já pode ser implentado -->
* Pointeiro de pilha: SP (2 bytes)
* Program counter: PC (2 bytes)
* Aritméticos/Lógicos: D0-D7 (8 registradores com 2 bytes cada)
<!-- * Outros: adicionar se necessário -->

<!-- ## Equipamentos de entrada/saída/armazenamento
* Entrada: leitor de arquivos externos (?)
* Saída: output no terminal (?)
* Armazenamento: memória (?) -->

## Blocos funcionais
* Unidade aritmética e lógica - Operações
* Unidade de controle - Instruções
<!-- * Sistema de Interrupções - Tipos: (?) -->

<!-- # Instruções -->

<!-- Definir o conjunto de instruções a ser implementado -->