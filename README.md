# Desafio 

Hoje, a ******, como dito anteriormente, é o maior ***** de ***** do Brasil.
Aqui trabalhamos constantemente com grande volume e complexidade de dados. Sabendo disso,
precisamos que você elabore uma solução que ofereça **armazenamento, processamento e disponibilização** desses dados, sempre considerando que tudo deve estar conforme as boas práticas de
segurança em TI. Afinal, nosso principal ativo são dados sensíveis dos consumidores brasileiros.

## Indice

* [Solução/arquitetura proposta]()
* [Tecnologias propostas]()
* [Payloads propostos]()
* [Um pouco da lógica]()
* [Outra ideia de solução]()


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
  * Garantir privacidade [GitLab](https://about.gitlab.com/)
  * Garantir disponibilidade [AWS](https://aws.amazon.com/pt/products/?nc2=h_ql_prod_fs_f)
