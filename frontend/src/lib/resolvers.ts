import gql from "graphql-tag";

// Resolvers receive the mutation call and run the according function.

const countQuery = gql`
  query Count {
    count @client
  }
`;

export const resolvers = {
  Mutation: {
    increment: (parent, args, { cache }) => {
      let data;
      try {
        data = cache.readQuery<any>({ query: countQuery });
      } catch {
        (err) => console.log(err);
      }
      cache.writeData({
        data: {
          count: data ? data.count + 1 : 1,
        },
      });
      return null;
    },
  },
};
