import gql from 'graphql-tag';

const typeDefs = gql `
    type appState {
        isDarkModeEnabled: Boolean
    }
`;

export default typeDefs;