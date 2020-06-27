import gql from 'graphql-tag'

export const ProfileDataQuery = gql`
    query Users($uuid: String!) {
      User (uuid: $uuid) {
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
