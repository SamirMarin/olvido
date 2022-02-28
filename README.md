# olvido
Keep track of your day to day

This is currently in development, initial prototype of scraper has been written, very much a prototype still :) 

Basic Idea is:

```mermaid
  graph TD;
      A[Scraper]-->B[mongoDB];
      A-->C[Notifier];
      C-->D[email,phone,slack,etc];
      E[Configuration_manager]-->B;
      F[GUI_CLIENT]-->E
      F-->A
```

