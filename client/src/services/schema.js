import gql from 'graphql-tag';

export const LOGIN = gql`
    mutation($input: LoginUser!) {
        loginUser(input: $input)
    }
`;

export const REFRESH_TOKEN = gql`
    mutation($input: RefreshTokenData!) {
        refreshToken(input: $input)
    }
`;

export const PARSE_TOKEN = gql`
    query {
        parseTokenData {
            username
            role
        }
    }
`;

