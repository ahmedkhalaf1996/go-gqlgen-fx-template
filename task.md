TASK
Implement an authenticated GraphQL API that allows users to track todos. The schema should include the mutations and queries as shown in the attached picture.

Authenticated access should be handled with the Authorization header and Bearer tokens. You can use the golang-jwt/jwt (https://pkg.go.dev/github.com/golang-jwt/jwt) package to generate tokens. If the user isn’t authenticated, the API should return a 401 using middleware/directives.

Your GraphQL schema should implement nested types and field resolvers or preloads to fetch them from the DB. This should work recursively.

ADDITIONAL IMPROVEMENTS (increase your chances against other candidates)
- Use the dataloader pattern to optimize N+1 queries.
- Use field collection to preload relationships & avoid N+1 altogether.
- Implement UUID with a custom scalar for todos.
- Add E2E tests with Ginkgo and Gomega.
- Use gorm/gen to generate typesafe repositories (https://gorm.io/gen/).

Commit your code to a VCS provider of your choice (GitHub, GitLab, Bitbucket, etc.) and share the link with us once completed. If you want to deploy the app, feel free to share the deployment link as well (see additional improvements).

Note that communication is only accepted via the before-mentioned e-mail. The E-Mail should contain the phrase blue elephant in white house.

Good luck!