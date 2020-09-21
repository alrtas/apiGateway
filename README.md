# Desafio 

Hoje, a ******, como dito anteriormente, é o maior ***** de ***** do Brasil.
Aqui trabalhamos constantemente com grande volume e complexidade de dados. Sabendo disso,
precisamos que você elabore uma solução que ofereça **armazenamento, processamento e disponibilização** desses dados, sempre considerando que tudo deve estar conforme as boas práticas de
segurança em TI. Afinal, nosso principal ativo são dados sensíveis dos consumidores brasileiros.

## Indice

* [Solução/arquitetura proposta](https://github.com/alrtas/apiGateway/blob/master/README.md#solu%C3%A7%C3%A3oarquitetura-proposta)
* [Tecnologias propostas](https://github.com/alrtas/apiGateway#tecnologias-propostas)
* [Endpoints & Payloads de entrada]
* [Endpoints & Payloads de saida]
* [Um pouco da lógica]()
* [Outra ideia de solução](https://github.com/alrtas/apiGateway#outra-ideia-de-solu%C3%A7%C3%A3o)


# Solução/arquitetura proposta

![](https://github.com/alrtas/apiGateway/blob/master/Utils/Imagens/arquitetura.PNG)

Proposta de criação de um middleware que faça a interface entre os microserviços / nanoserviços com as bases de dados e outros microserviços internos, tornando trasnparente para a aplicação qual base de dados será acessado e qual tecnologia/arquitetura estão estruturada (ou não) os dados.
 Sendo assim haverá uma padronização na inclusão e disponibilização dos dados. E os micro/nano serviços deveram apenas se preocupar em integrar via API com o middleware e não diretamente com as bases de dados, que estão em uma [VPC](https://aws.amazon.com/pt/vpc/), garantindo que a unica maneira de acessar os dados é através o consumo das APIs conforme vide no readme

### Vantagens 
* Segurança, só é necessário cria o permissionamento uma vez, entre HOST e middleware. E não com cada base de dados.
* Segurança, as bases de dados não estão disponiveis na internet.
* Se tiver adotado uma arquitetura de microsserviços, uma única solicitação pode exigir chamadas para dezenas de aplicações distintas.
* Arquitetura, a alteração no tipo ou tencnologia dos bancos de dados não causa alterações na comunicação com as aplicações.
* Com o tempo, você incluirá novos serviços de API e descontinuará outros, mas seus clientes ainda vão querer encontrar todos os seus serviços no mesmo lugar.
* Você quer entender como as pessoas usam suas APIs, então inclui ferramentas de monitoramento e análise.

### Desvantagens
* Há uma curva de aprendizado quando se trata de arquitetar aplicativos de alta disponibilidade em escala, especialmente porque o gateway de API será o único ponto de entrada entre o front-end e as APIs e também atuará como um único ponto de falha.

## Tecnologias propostas

* MiddleWare / API Gateway utilizaremos o [GOLang](http://golang.org/)
* Base A utilizaremos [MySQL](https://www.mysql.com/)
  * Maiores níveis de segurança, mas o acesso a esses dados não é tão performático
* Base B utilizaremos [ElasticSearch](https://www.elastic.co/pt/)
  * Acesso mais rápido/performatico.
* Base C utilizaremos [Redis](https://redis.io/)
  * Acesso extremamente rápido.
* AWS + Jenkins + GitLab
  * Garantir CI/CD [Jenkins](https://www.jenkins.io/)
  * Garantir privacidade e versionamento [GitLab](https://about.gitlab.com/)
  * Garantir disponibilidade, Dockers na [AWS](https://aws.amazon.com/pt/products/?nc2=h_ql_prod_fs_f)
  
## Endpoints & Payloads de entrada

* URI: `api.tenant.com/cadastro/geral`
 * Authorization : `your token`
 * Verbo: `POST`
 * Payload da requisição : [disponivel aqui]()
 
 <br />

* URI: `api.tenant.com/cadastro/financeiro`
 * Authorization : `your token`
 * Verbo: `POST`
 * Payload da requisição : [disponivel aqui]()

<br />

* URI: `api.tenant.com/cadastro/dividas`
 * Authorization : `your token`
 * Verbo: `POST`
 * Payload da requisição : [disponivel aqui]()
   
<br />
  
## Endpoints & Payloads de saida


## Um pouco da lógica

## Outra ideia de solução
![](https://github.com/alrtas/apiGateway/blob/master/Utils/Imagens/wso2_enterprise_integrator.PNG)


Outra alternativa é utilizar o [WSO2](https://wso2.com/integration/) como API Manager/gateway e integrador entre diferentes soluções ou [ESB](https://wso2.com/products/enterprise-service-bus/), neste casos as integrações serão feitas somente com base de dados, e disponibilizadas e consumidas via layer implementado pelo API manager. Aumentando a velocidade e agilidade no desenvolvimento de futuras integrações, sem necessidade de criar arquitetura, sem codigo, somente focando em entregar valor para o negócio.

Segue abaixo um video da plataforma WSO2.
[LINK](https://www.youtube.com/watch?v=hs_FLM5a6Ck&list=PLp0TUr0bmhX6PMqphqe6dJiap3B5KpLyi&ab_channel=WSO2)



### Payload de cadastro geral
    {
	"resource":{
		"nome":"Thiago Alberto da silva",
		"cpf":"09489601918",
		"telefoneCelular":"48996260373",
		"telefoneResidencial":"4833745517",
		"email":"alrrtas@gmail.com",
		"emailsecundario":"thiagos.tas@gmail.com",
		"nascimento":"16/08/1997",
		"sexo" : "masculino",
		"estadoCivil" : "solteiro",
		"enderecos": [
						{
							"tipo":"residencial"
							"endereco": "Rua Arcanjo Cunha",
							"bairro": "Rio Grande",
							"cidade": "Palhoça",
							"numero": 56,
							"referencia" : "ao lado da padaria",
							"uf": "SC",
							"cep": "88131700",
							"ibge": "4211900"
						}
					]
	}
    }
 
### Payload de cadastro financeiro
    {
	"resource": {
		"cpf": "09489601918",
		"rendaMensal" : "6000.00"
		"profissao" : "analista de sistemas",
		"patrimonio" : "150000.00",
		"quantidadeDeBens" : "3",
		"bens" : [
					{
						"tipo":"imovel",
						"estado" : "pago",
						"valor" : "100000.00"
					},
					{
						"tipo":"automovel",
						"estado":"financiado"
						"valor":"50000.00",
						"pago" : "25000.00",
						"debitos" : "25000.00"
					},
					{
						"tipo": "investimento",
						"valor" : "25000.00"
					}
					
		]
	}
    }
### Payload de cadastro de dividas
    {
	"resource": {
		"cpf": "09489601918",
		"quantidadeDeDividas":"3",
		"dividasEmAtraso":"1",
		"valorTotalDevido":"25000.00",
		"dividas" : [
						{
							"tipo":"financiamento",
							"valor":"10000.00",
							"status": "em dia"
						},
						{
							"tipo":"cartao",
							"valor":"5000.00",
							"status": "atrasado"
						},
						{
							"tipo":"financiamento",
							"valor":"10000.00",
							"status": "em dia"
						}
		]
	}
    }
