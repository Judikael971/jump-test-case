# Usages

## Lancement du docker
Depuis la racine du projet (au niveau du `Makefile`) exécuter la commande :

```sh
make usage
```
Qui vous détaillera les raccourcis possibles pour utiliser le container.

### Lancer le container
```sh
make up
```
La commande va créer et démarrer le container

### Arrêter le container
```sh
make down
```
La commande va arrêter, supprimer le container puis supprimer l'image.

## Requêter l'API
J'ai mis à disposition la [collection postman](../api/jump.postman_collection.json) pour appeler les routes sans lignes de commandes.

Sinon, vous trouverez les requêtes curl à exécuter dans votre invite de commande.

### Route 'users'
Appel de la route 'users' pour récupérer la liste des utilisateurs.

Commande :

```sh
curl http://localhost:8999/users --request "GET" --include
```

### Route 'invoice'
Appel de la route 'invoice' pour créer une facture.

Commande :

```sh
curl http://localhost:8999/invoice \
 --include \
 --header "Content-Type: application/json" \
 --request "POST" \
 --data '{"user_id":24,"amount":113.45,"label":"Work for April"}'
```

*Les valeurs du paramètre `data` de la requête `curl` sont données à titre d'exemple.
Veuillez les remplacer par de véritable valeur.*

### Route 'transaction'
Appel de la route 'transaction' permettant d'ajouter une transaction monétaire pour une facture donnée.

Commande :

```sh
curl http://localhost:8999/transaction \
 --include \
 --header "Content-Type: application/json" \
 --request "POST" \
 --data '{"invoice_id":42,"amount":956.32,"reference":"JMPINV200220117"}'
```

*Les valeurs du paramètre `data` de la requête `curl` sont données à titre d'exemple.
Veuillez les remplacer par de véritable valeur.*
