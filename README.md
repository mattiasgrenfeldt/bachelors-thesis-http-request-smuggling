# Bachelor's thesis on HTTP Request Smuggling

During the spring of 2021, we (Mattias Grenfeldt and Asta Olofsson) wrote our bachelor's thesis in Computer Science at KTH Royal Institute of Technology in Sweden. We studied HTTP Request Smuggling. The thesis can be found [here](https://urn.kb.se/resolve?urn=urn:nbn:se:kth:diva-302371).

Once all issues found have been fixed by the affected vendors, we will publish here the code used for the test harness, the test requests and which systems are which in the anonymized report.

## IEEE EDOC 2021 paper

The bachelor's thesis was later rewritten into a conference paper with the help of Viktor Engström, and Robert Lagerström. The paper was submitted to [IEEE EDOC 2021](https://ieee-edoc.org/2021/) and got accepted.

## Systems investigated

Here is the de-anonymization of the systems we investigated:

- P3 - ???
- P4 - ???
- P2 - ???
- P5 - ???
- P1 - ???
- P6 - ???
- S1 - [Gunicorn](https://gunicorn.org/)
- S2 - ???
- S3 - ???
- S4 - ???
- S5 - ???
- S6 - ???

## Errata

After the thesis was published, we realized that we had interpreted the situation with `Transfer-Encoding: chunked` and HTTP version 1.0 incorrectly. It was very unclear what a correct interpretation was. So we opened an issue on the specification. [Here](https://github.com/httpwg/http-core/issues/879) is the discussion that followed. This resulted in a change in the specification.
