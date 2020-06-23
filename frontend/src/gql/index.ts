import gql from 'graphql-tag'

export const ProfileDataQuery = gql`
    query Users($uuid: String!) {
      Users {
        uuid
        username
        email
        description
        profile_picture
      }
      followers: Follows(followee: $uuid) {
        uuid
      }
      following: Follows(follower: $uuid) {
        uuid
      }
    }
`;
