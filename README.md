# gocars 🚗

gocars is een multi-tenant webapplicatie waarmee autobedrijven eenvoudig hun eigen website en voorraad kunnen beheren.
De applicatie is gebouwd in Go, gebruikt server-side templates en htmx voor een snelle SPA-achtige ervaring zonder zware frontend frameworks.

## Kernidee

- Er is één Go-applicatie
- Meerdere autobedrijven (tenants) gebruiken dezelfde app
- Elk autobedrijf heeft:
    - een eigen website
    - een eigen dashboard
    - volledig gescheiden data (auto’s, teksten, gebruikers)

## Multi-tenant structuur

Elke request hoort bij één tenant (autobedrijf).
De tenant wordt bepaald op basis van het domein of subdomein, bijvoorbeeld:

```
jans.gocars.nl
piet.gocars.nl
```

Middleware leest het domein en koppelt de request aan het juiste autobedrijf (dealer).
