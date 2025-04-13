# Desafio PROGRAMAÇÃO - Sistema de Cadastros em Go 🚀

### Desafio criado por Lucas Carrilho - *@devmagro* ([LinkedIn](https://www.linkedin.com/in/karilho/))

## Introdução
Você chegou ao **primeiro grande desafio do guia**! 😎 Agora, é hora de colocar a mão na massa e aplicar seus conhecimentos na prática, criando um **Sistema de Adoção para Pets em Go!**

Leia TODO o desafio com ATENÇÃO antes de começar e respeite as regras. Evite o uso de IA generativa ao máximo e use-a com sabedoria apenas quando necessário!

## Objetivo
Neste desafio, você criará um sistema de cadastro via CLI (interface de linha de comando) em Go, permitindo que futuros adotantes escolham seus pets. O dono do abrigo de animais, usuário do programa, poderá:

- Cadastrar um novo pet
- Buscar dados de um pet cadastrado
- Alterar os dados de um pet cadastrado
- Deletar um pet cadastrado
- Listar todos os pets cadastrados
- Listar pets por critérios (idade, nome, raça)
- E mais!

O sistema deve ser implementado utilizando structs, manipulação de arquivos, organização em pacotes e boas práticas de Go. 🚀

## Conhecimentos que você irá aplicar:
- Structs e métodos em Go
- Manipulação de arquivos e arrays/slices
- Tratamento de erros em Go
- Boas práticas de código
- Organização de pacotes

# Requisitos do Desafio 📋

O desafio é dividido em **PASSOS**, nos quais você desenvolverá novas funcionalidades para a aplicação CLI. Sinta-se à vontade para criar structs, funções e pacotes conforme achar necessário!

Para começar, faça um *fork* deste repositório, clone-o em sua máquina local e deixe sua ⭐ no repositório.

## 📍 Passo 1 - Leitura do arquivo de base com as perguntas essenciais 📄
Crie um arquivo chamado `formulario.txt` com as seguintes perguntas:

**1 - Qual o nome e sobrenome do pet?**  
**2 - Qual o tipo do pet (Cachorro/Gato)?**  
**3 - Qual o sexo do animal?**  
**4 - Qual endereço e bairro que ele foi encontrado?**  
**5 - Qual a idade aproximada do pet?**  
**6 - Qual o peso aproximado do pet?**  
**7 - Qual a raça do pet?**

Sua aplicação deve ler esse arquivo `.txt` e exibir as perguntas no terminal.

**<span style="color:red">Não é permitido usar `fmt.Println` para exibir as perguntas diretamente no código; você deve LER o arquivo!</span>**

## 📍 Passo 2: Criação do menu inicial 📝
Após exibir o conteúdo do `formulario.txt`, crie um menu inicial com as opções:

1. Cadastrar um novo pet
2. Alterar os dados do pet cadastrado
3. Deletar um pet cadastrado
4. Listar todos os pets cadastrados
5. Listar pets por algum critério (idade, nome, raça)
6. Sair

O menu deve aparecer no terminal, e o usuário deve escolher uma opção digitando o número correspondente.

### _Regras ⚠️_
1. O usuário não pode usar 0 ou números negativos.
2. Se o usuário digitar um número inválido, o menu deve ser exibido novamente.
3. Só números são permitidos (sem letras ou caracteres especiais).

## 📍 Passo 3: Cadastro de novos Pets 📝
Ao selecionar a opção 1, o usuário deve cadastrar um novo pet. Leia as perguntas do `formulario.txt` e armazene as respostas em uma struct `Pet`.

### _Regras ⚠️_
1. O nome e sobrenome do pet são obrigatórios; caso contrário, retorne um erro.
2. O nome completo só pode conter letras de A-Z (sem caracteres especiais).
3. Para o endereço, pergunte:
   - Número da casa
   - Cidade
   - Rua
4. Idade e peso aceitam números com vírgulas ou pontos, mas apenas números.
5. Peso maior que 60kg ou menor que 0.5kg gera erro.
6. Idade maior que 20 anos gera erro.
7. Idade menor que 1 ano (em meses) deve ser convertida para 0.X anos.
8. Raça não pode conter números ou caracteres especiais.
9. Se nome, raça, peso, idade ou número do endereço não forem informados, preencha com a constante `"NÃO INFORMADO"`.

## 📍 Passo 4: Armazenamento em Arquivo TXT 📂
Salve as respostas em um arquivo `.txt` com o nome no formato: `20231101T1234-FLORZINHADASILVA.txt`.

### _Regras ⚠️_
1. O nome do arquivo deve seguir o formato: ano, mês, dia, "T", hora, minuto - NOME+SOBRENOME em maiúsculas.
2. Salve o arquivo na pasta `petsCadastrados` na raiz do projeto.
3. O conteúdo deve ter cada resposta em uma linha, sem as perguntas. Exemplo:
   ```
   Florzinha da Silva
   Gato
   Femea
   Rua 2, 456, Seilandia
   6 anos
   5kg
   Siames
   ```

## 📍 Passo 5: Buscar os dados do Pet Cadastrado 🔍
Na opção 2, o usuário deve buscar pets por critérios como:
- Nome ou sobrenome
- Sexo
- Idade
- Peso
- Raça
- Endereço

Permita combinar 1 ou 2 critérios (ex.: nome e idade). Exiba um submenu para escolher os critérios e mostre os resultados no formato:

```
1. Rex - Cachorro - Macho - Rua 1, 123 - Cidade 1 - 2 anos - 5kg - Vira-lata
2. Florzinha da Silva - Gato - Femea - Rua 2, 456 - Seilandia - 6 anos - 5kg - Siames
```

### _Regras ⚠️_
1. O critério "tipo de animal" é obrigatório e deve ser escolhido primeiro.
2. A busca por nome deve incluir partes (ex.: "FLOR" encontra "Florzinha").
3. A busca deve ser *case-insensitive* (ignorar maiúsculas/minúsculas e acentos).

### _Regra nível 2 (opcional) ⚠️_
1. Permita busca por data de cadastro (mês e ano) combinada com outros critérios.
2. Destaque o termo pesquisado nos resultados (ex.: negrito). Use códigos ANSI para formatação no terminal.

## 📍 Passo 6: Alterar um pet cadastrado ✍️
Na opção 3, permita alterar os dados de um pet.

### _Regras ⚠️_
1. Use os critérios de busca do Passo 5.
2. Exiba a lista de resultados e peça o número do pet a alterar.
3. Permita alterar todos os dados, exceto `tipo` e `sexo`.
4. Se o número for inválido, repita o menu de busca.
5. Após alterar, salve o arquivo `.txt` com as novas informações.

## 📍 Passo 7: Deletar um animal cadastrado 🔍
Na opção 4, permita deletar um pet.

### _Regras ⚠️_
1. Use os critérios de busca do Passo 5.
2. Exiba a lista e peça o número do pet a deletar.
3. Solicite confirmação ("SIM" ou "NÃO").
4. Delete o arquivo `.txt` correspondente se "SIM".

## 📍 Passo 8: Sair
Na opção 7, encerre o programa.

## 📍 Passo EXTRA ☠️☠️☠️
Adicione um menu inicial extra:
```
1 - Iniciar o sistema para cadastro de PETS
2 - Iniciar o sistema para alterar formulário
```

Na opção 2, exiba:
1. Criar nova pergunta
2. Alterar pergunta existente
3. Excluir pergunta existente
4. Voltar ao menu inicial
5. Sair

### _Regras Opção 1 ⚠️_
1. Adicione a nova pergunta ao `formulario.txt` com numeração sequencial.
2. Respostas são opcionais; use "NÃO INFORMADO" se vazio.

### _Regras Opção 2 ⚠️_
1. Só altere perguntas além das 7 originais.
2. Peça o número da pergunta, exiba-a e permita editá-la.

### _Regras Opção 3 ⚠️_
1. Só delete perguntas além das 7 originais.
2. Confirme com "SIM" ou "NÃO".

### _Regras Gerais ⚠️_
1. Reorganize a numeração se uma pergunta for deletada.
2. Evite linhas vazias no `formulario.txt`.
3. Ao cadastrar um pet, inclua as perguntas extras e salve as respostas no arquivo `.txt` do pet, com o formato:
   ```
   1 - Florzinha da Silva
   ...
   7 - Siames
   8 - [EXTRA - PERGUNTA NOVA ADICIONADA] - RESPOSTA DO USUÁRIO
   9 - [EXTRA - PERGUNTA NOVA ADICIONADA] - RESPOSTA DO USUÁRIO
   ```

## Considerações Finais 📝
- Faça *commits* conforme avança, seguindo o padrão [Conventional Commits](https://www.youtube.com/watch?v=sbK9h45Jc5o).
- Crie um README.md com a descrição do projeto e instruções de execução.

### Gostou do projeto?
Parabéns por chegar até aqui! Você concluiu o primeiro grande desafio adaptado para Go. 🚀