# DOC
## Sécurités / Vulnérabilités
La commande `docker scout quickview` relève une vulnérabilité sur l'image `postgres:14.2-alpine`. 

Afin de se protéger, j'ai passé postgres à la version `postgres:15.4-alpine`.

Il reste une vulnérabilité sur l'image docker que j'ai créé ([CVE-2023-3978](https://nvd.nist.gov/vuln/detail/CVE-2023-3978)), pour l'instant elle n'est pas encore patchée.

## Schema DB
J'ai mis à jour votre [schema.sql](../build/package/schema.sql).
### Modification de l'existant
En changeant le type de deux colonnes :
- *amount* (`BIGINT`) de la table **invoices** par `FLOAT`
- *balance* (`BIGINT`) de la table **users** par `FLOAT`

Permet d'enregistrer des valeurs décimales.

### Ajout de tables
Afin de conserver une trace des transactions, j'ai ajouté deux tables :
- Une qui contiendra la liste des états de transactions `transaction_status`. 
- Une pour stocker les demandes de transactions dont les factures sont connues `transactions`.

## Configuration des LOGS
### Log HTTP
Enregistre sous forme de JSON les requêtes que reçoit l'application.

### Log SQL (GORM)
Enregistre toutes les requêtes SQL que lancera l'application (non formalisé en JSON, mais c'est possible via middleware).

On peut changer le niveau de criticité déclenchant l'enregistrement du log.