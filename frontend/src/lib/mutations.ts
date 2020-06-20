import gql from "graphql-tag";

// This is a mutation, which is called by the returned varioble from useMutation() in a component.
// Then apollo client receives it and finds the approprate resolver.

export const INCREMENT = gql`
  mutation Increment {
    increment @client
  }
`;
