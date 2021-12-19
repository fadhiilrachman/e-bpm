import React, { useEffect } from 'react'
import { navigate } from 'gatsby'
import { createTheme, ThemeProvider } from '@mui/material/styles'

import AuthLoginForm from '../../components/auth/login'

import { ApolloProvider } from '@apollo/react-hooks'
import client from '../../services/graphql'

const theme = createTheme()

export default () => {
    useEffect(() => {
        const token = localStorage.getItem('token')
        if (token) {
            navigate('/')
        }
    }, [])

    return (
        <ThemeProvider theme={theme}>
            <ApolloProvider client={client}>
                <AuthLoginForm />
            </ApolloProvider>
        </ThemeProvider>
    );
};