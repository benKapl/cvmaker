# CVMAKER
![Go](https://img.shields.io/badge/Go-1B73BA?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white)
![Scaleway](https://img.shields.io/badge/Scaleway-000000?style=for-the-badge&logo=scaleway&logoColor=FFFFFF)
![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white)
![React](https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)
![Remix](https://img.shields.io/badge/Remix-181818?style=for-the-badge&logo=remix&logoColor=FFFFFF)

> En cours de développement

Générateur de CV automatisé permettant de créer un CV d'une page parfaitement adapté à une offre d'emploi et le profil de l'utilisateur.

## Stack
### API
- Go avec la librairie net/http pour le serveur API
- Base de données Postgres, utilisation de Goose pour la migration et SQLC pour la génération de code Go permettant d'intéragir avec la base de données.
- Services tiers : modèles de Mistral (Ollama ou API) pour la génération de chaque élément du CV
  => Mise à profit de la concurrence native de Go: chaque partie du CV (Formation, Expériences, etc.) est générée en parallèle par l'exécution d'une goroutine.

### Web
- Langage : React TypeScript
- Framework : React Router v7 (anciennement Remix)

### Build
- Docker
- Scaleway pour l'hébegement de l'application :
  - Scaleway Container Registry pour l'enregistrement des images Docker
  - Scaleway Managed Postgres pour la gestion de la base de données
  - Scaleway Container Run pour faire tourner le serveur API
  - Scaleway Object Storage pour l'hébergement du front (site statique)
- Github Actions pour la pipeline automatisée de CI-CD      

## Roadmap
- Authentification
- Import des données nécessaires à la caractérisation du/de la candidate
- Génération de CV sur la base d'un template
- Téléchargement PDF du CV
- Modification des informations importées / générées
- Création de nouveaux templates
- Limitation du nombre de génération (sur la base de crédits)








