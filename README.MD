# SPORTIN-API

The Sportin' API is an API REST made for the athletes, who need to track down every step they make and want to personalize their workout.

This API is realised by :
- *Sorgiati Sacha*
- *Laurent Rémi*
- *Lemée Etienne*
- *Bamas Lunay Rémy*

## Installation

#### Prerequisites

- Go 1.23
- Database : MariaDB
 
#### Steps

1. Clone the repository

```bash
git clone https://github.com/your-username/strengthapi.git
cd strengthapi
```

2. Copy/paste env.example rename .env and set your env variables.

3. install depedencies

```bash
go mod tidy
```

4. Launch the project

```bash
go run main.go
```

## Features

#### Users

1. Setup an account with your email adress and your password
1. Follow your stats like your weight, calories etc

#### Exercises and programs

1. Create some new exercises or follow the exercises already existing
1. Add these exercises in programs for your workouts.
1. These exercises are listed in categories wich can correspond to the muscles worked on or the workout type they can be used in.
1. These exercices can also be intensified.
1. For your programs, your can setup some sets to make your workout complete.


## Endpoints

For the listing of all our endpoints, we used **Swagger** so you can find every route and what is it used for more easily.


**Format :**
- @Summary Method name
- @Description What is it used for
- @Tags Which group is it on
- @Accept Payload format that can be sent
- @Produce Payload format that can be received
- @Param prerequisites that need to be present to access correctly to this route
- @Success or @Failure statut code and what's sent
- @Router Endpoint where you can find it

You can find under every tag which method is allowed for what you're trying to get on the API (users, exercises, programs, ...).

### **Thanks for using the Sportin API**