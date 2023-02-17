const { ApolloServer } = require("apollo-server");
const { ApolloGateway } = require("@apollo/gateway");

const gateway = new ApolloGateway({
  serviceList: [
    { name: "organization", url: process.env.ORGANIZATION_GRAPHQL_API },
  ],
});

const server = new ApolloServer({
  gateway,
  subscriptions: false,
});

server
  .listen({
    port: process.env.PORT,
    host: "0.0.0.0",
  })
  .then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
  });
