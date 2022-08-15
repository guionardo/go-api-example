# go-api-example

## Proposta
Desenvolver uma API que exponha os dados disponíveis em [zip](http://www.prefeitura.sp.gov.br/cidade/secretarias/upload/chamadas/feiras_livres_1429113213.zip) utilizando uma abordagem orientada a recursos e que atenda os requisitos listados abaixo.

## Escopo

Utilizando os dados do arquivo “DEINFO_AB_FEIRASLIVRES_2014.csv”, implemente:

* [X] cadastro de uma nova feira;
* [X] exclusão de uma feira através de seu código de registro;
* [X] alteração dos campos cadastrados de uma feira, exceto seu código de registro;
* [X] busca de feiras utilizando ao menos um dos parâmetros abaixo:
    * [X] distrito
    * [X] regiao5
    * [X] nome_feira
    * [X] bairro
  
## Requisitos

* [X] utilize git ou hg para fazer o controle de versão da solução do teste e hospede-a no Github ou Bitbucket;
* [X] armazene os dados fornecidos pela Prefeitura de São Paulo em um banco de dados relacional que você julgar apropriado;
* [X] a solução deve conter um script para importar os dados do arquivo “DEINFO_AB_FEIRASLIVRES_2014.csv” para o banco relacional;
* [X] a API deve seguir os conceitos REST;
* [X] o Content-Type das respostas da API deve ser “application/json";
* [ ] o código da solução deve conter testes e algum mecanismo documentado para gerar a informação de cobertura dos testes;
* [ ] a aplicação deve gravar logs estruturados em arquivos texto;
* [ ] a solução desta avaliação deve estar documentada em português ou inglês. Escolha um idioma em que você seja fluente;
* [ ] a documentação da solução do teste deve incluir como rodar o projeto e exemplos de requisições e suas possíveis respostas;
