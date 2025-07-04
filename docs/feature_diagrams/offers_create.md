```mermaid
sequenceDiagram
  actor User as User
  participant Web as React Frontend
  participant API as Go HTTP Router (mux)
  participant Handlers as handlerOffersCreate
  participant OfferService as OfferService
  participant Prompter as Prompter
  participant LLMClient as LLM Client (AI)
  participant DBPackage as Database Package
  participant PostgreSQL as PostgreSQL

  User ->>+ Web: Pastes job offer text, clicks Create
  Web ->>+ API: POST /api/offers (offer: raw_text)
  API ->>+ Handlers: Routes to handlerOffersCreate
  Handlers ->> Handlers: Extracts userID from JWT context
  Handlers ->>+ OfferService: Calls CreateOffer(userID, rawOffer)
  OfferService ->>+ Prompter: Calls ParseOffer(rawOffer)
  Prompter ->>+ LLMClient: Sends structured prompt with job offer text
  LLMClient ->>+ LLMClient: AI processes and extracts structured data
  
  alt AI parsing successful
    LLMClient -->>- Prompter: Returns JSON (title, organization, missions, stack, etc.)
    Prompter ->> Prompter: Validates and handles missing required fields
    Prompter -->>- OfferService: Returns ParsedOffer struct
    OfferService ->> OfferService: Maps ParsedOffer to database params
    OfferService ->>+ DBPackage: CreateOffer(structured_data)
    DBPackage ->>+ PostgreSQL: INSERT INTO offers (title, organization, missions, ...)
    PostgreSQL -->>- DBPackage: Returns new offer record
    DBPackage -->>- OfferService: (Offer, nil)
    OfferService -->>- Handlers: (Offer, nil)
    Handlers ->> Handlers: Converts DB offer to API response format
    Handlers -->>- API: 201 Created (structured offer data)
    API -->>- Web: 201 Created (parsed offer with extracted fields)
    Web -->>- User: Shows parsed offer details
  else AI parsing failed
    LLMClient -->>- Prompter: Returns error
    Prompter -->>- OfferService: (nil, parsing error)
    OfferService -->>- Handlers: (nil, parsing error)
    Handlers -->>- API: 500 Internal Server Error
    API -->>- Web: 500 Internal Server Error
    Web -->>- User: Shows "Failed to parse offer" error
  end
```