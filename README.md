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

---

## Hosts file (voor lokale tenants)

Voeg tenant-subdomains toe aan je hosts-bestand.

### Windows

1. Open **Notepad als administrator**
2. Open:

   ```
   C:\Windows\System32\drivers\etc\hosts
   ```
3. Voeg toe:

   ```
   127.0.0.1 tenant1.localhost
   127.0.0.1 tenant2.localhost
   ```

### Linux / macOS

```bash
sudo nano /etc/hosts
```

Voeg toe:

```
127.0.0.1 tenant1.localhost
127.0.0.1 tenant2.localhost
```

Daarna:

* [http://tenant1.localhost:8080](http://tenant1.localhost:8080)
* [http://tenant2.localhost:8080](http://tenant2.localhost:8080)

---
