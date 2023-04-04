Disciplina
Vamos repassar o ciclo novamente:
 * Escrever um teste
 * Compilar o código sem erros
 * Rodar o teste, ver o teste falhar e certificar que a mensagem de erro faz sentido
 * Escrever a quantidade mínima de código para o teste passar
 * Refatorar

O processo TDD e por que as etapas são importantes
 * Escreva um teste que falhe e veja-o falhar, para que saibamos que escrevemos um teste relevante para nossos
   requisitos e vimos que ele produz uma descrição da falha fácil de entender
 * Escrever a menor quantidade de código para fazer o teste passar, para que saibamos que temos software funcionando
 * Em seguida, refatorar, tendo a segurança de nossos testes para garantir que tenhamos um código bem feito e fácil de 
   trabalhar

No nosso caso, passamos de Ola() para Ola("nome"), para Ola ("nome"," Francês ") em etapas pequenas e fáceis de entender.

Naturalmente, isso é trivial comparado ao software do "mundo real", mas os princípios ainda permanecem. O TDD é uma 
habilidade que precisa de prática para se desenvolver. No entanto, você será muito mais facilidade em escrever software
conseguindo dividir os problemas em pedaços menores que possa testar.
