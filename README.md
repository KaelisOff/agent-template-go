# Modèle d'agent Kaelis (Go)

Point de départ pour écrire un **agent** pour la box Kaelis : une automatisation qui tourne sur la box,
avec son modèle d'IA, sans réseau et sans accès aux données, en passant uniquement par les capacités
que le propriétaire a acceptées.

Ce modèle contient un agent qui marche déjà (un mot d'encouragement chaque matin de semaine), ses
tests, et tout ce qu'il faut pour construire un paquet `.kagent` reproductible. Il suffit de Go 1.24
ou plus récent : ni Docker, ni box pour développer.

## Démarrer

1. **Use this template** → *Create a new repository*. Le dépôt doit être **public** : les agents sont
   open source, et le magasin les reconstruira depuis leurs sources.
2. Dans [`manifest.toml`](manifest.toml), remplacez :
   - `id` : `io.github.<votre-compte>.<nom-de-l-agent>`, en minuscules ;
   - `name`, `description`, `author`, `source` (l'adresse de votre dépôt).
3. Écrivez votre logique dans [`main.go`](main.go), et ne demandez dans `[permissions]` que les
   capacités dont l'agent a besoin : le propriétaire les verra toutes avant d'installer.
4. Vérifiez :

```sh
make test      # tests, avec une box simulée
make package   # construit dist/<id>-<version>.kagent
make verify    # reconstruit tout et vérifie que le paquet est identique
```

## Essayer sur une box

Le propriétaire ou un administrateur de la box installe le paquet depuis le **Centre d'agents**
(*Installer un agent*), après avoir lu et accepté les permissions. *Lancer maintenant* fait un
passage tout de suite ; l'*Activité* de l'agent montre le résumé et le journal de chaque passage.

## Publier une version

Changez `version` dans le manifeste, puis :

```sh
git tag v0.2.0
git push origin v0.2.0
```

La CI vérifie que le tag correspond à la version du manifeste, lance les tests, construit le paquet
deux fois pour vérifier qu'il est reproductible, et le joint à une *release* GitHub.

## Pour aller plus loin

- [Le kit de développement](https://github.com/KaelisOff/agent-sdk-go) : capacités, tests avec
  `agenttest`, conseils pour écrire un agent ;
- [le manifeste](https://github.com/KaelisOff/agent-sdk-go/blob/main/docs/manifeste.md),
  [le protocole](https://github.com/KaelisOff/agent-sdk-go/blob/main/docs/protocole.md).

## Licence

Ce modèle est sous [licence Apache 2.0](LICENSE). C'est la licence de votre agent si vous la gardez ;
vous pouvez en choisir une autre, du moment qu'elle est libre (pensez à changer `license` dans le
manifeste et le fichier `LICENSE`).
