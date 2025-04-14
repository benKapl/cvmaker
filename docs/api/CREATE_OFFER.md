## **Create offer**

Format a job offer to cvmaker's format and save it to database

- **URL**

  /offers

- **Method:**

  `POST`

- **Data Params**

  `body=[text]`

- **Success Response:**

  - **Code:** 201 <br />
    **Content:**
    ```json
    {
      "id": 12,
      "created_at": "2018-12-10T13:45.000Z",
      "updated_at": "2018-12-10T13:45.000Z",
      "label": "Backend Developer GO (F/N/M)",
      "organization": "Leboncoin Tech",
      "organization_description": "Created in 2006, leboncoin.fr is an exchange platform focused on simplifying consumption and fostering local relations through digital tools. It is the leading site for private sales and the 5th most visited site in France, employing 1,400 people.",
      "missions": "To join the 'Maps' feature team and contribute to revolutionizing the local ad search experience by developing cutting-edge mapping features. This involves redefining cartographic search standards, enhancing geographical precision, and developing algorithms to highlight relevant properties.",
      "stack": [
        "Golang (Go)",
        "Backend Development",
        "Microservices",
        "Distributed Architectures",
        "Event-Driven Systems",
        "RESTful APIs",
        "PostgreSQL",
        "Kafka",
        "Message Brokers",
        "Agile",
        "Scrum",
        "Kanban Flow",
        "Mapping",
        "Geospatial Search"
      ],
      "expected_profile": [
        "At least 5 years of backend development experience",
        "Minimum 2 years experience with Golang",
        "Experience with mapping/geospatial search topics",
        "Experience with distributed architectures, event-driven systems, microservices, RESTful APIs",
        "Familiarity with Kafka",
        "Proficiency in PostgreSQL",
        "Ability to evaluate architectural decisions, make trade-offs, implement them",
        "Knowledge and practice of Agile methodologies",
        "Aptitude for teamwork, knowledge sharing, helping others",
        "Good level of English for technical communication"
      ],
      "miscellaneous": [
        "Location: Paris (10th arrondissement), France",
        "Work Arrangement: Hybrid (Partial remote, min 2 days on-site per 2-week sprint)",
        "Job Type: Full-time, 'Cadre' status",
        "Team Context: Dynamic, multidisciplinary, international 'Maps' team",
        "Interview Process: HR phone screen, Tech test + interview, Manager interview, Optional N+2 interview",
        "Benefits include: Work From Anywhere (up to 20 days/year), Employee Assistance Program",
        "Company Culture: Values diversity, encourages applying even if not all requirements met"
      ]
    }
    ```

- **Error Response:**

  <_Most endpoints will have many ways they can fail. From unauthorized access, to wrongful parameters etc. All of those should be liste d here. It might seem repetitive, but it helps prevent assumptions from being made where they should be._>

  - **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "Log in" }`

  OR

  - **Code:** 422 UNPROCESSABLE ENTRY <br />
    **Content:** `{ error : "Email Invalid" }`

- **Sample Call:**

  <_Just a sample call to your endpoint in a runnable format ($.ajax call or a curl request) - this makes life easier and more predictable._>

- **Notes:**

  <_This is where all uncertainties, commentary, discussion etc. can go. I recommend timestamping and identifying oneself when leaving comments here._>
