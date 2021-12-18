import { ApolloClient, InMemoryCache } from '@apollo/client';
import fetch from 'isomorphic-fetch';

const client = new ApolloClient({
    fetch,
    uri: 'http://localhost:8080/query',
    cache: new InMemoryCache(),
    request: (opr) => {
        const token = localStorage.getItem('token');
        opr.setContext({
            headers: {
                authorization: token ? `Bearer ${token}` : '',
            },
        });
    },
});

export default client;