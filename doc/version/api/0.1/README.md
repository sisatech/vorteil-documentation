# API Changelogs

Changelogs, feature pipelines, and other information related to the API's 
version.

## 0.1.2 (2018-11-05)
- vcfgLinter query will no longer return false warnings if a toml parsing error occurs.
- The 'query' argument will is now correctly used to filter results on supported 'list' queries (such as 'listBuckets', and 'listListeners').

## 0.1.1 (2018-09-26)
- Fixed an issue with the 'importSharedObjects' GraphQL mutation, which placed some shared objects in incorrect top-level directories within the target project.