import React, { useEffect } from 'react'
import { navigate } from 'gatsby'
import { createTheme, ThemeProvider } from '@mui/material/styles'

const pageStyles = {
    color: '#232129',
    padding: '96px',
    fontFamily: '-apple-system, Roboto, sans-serif, serif',
}
const headingStyles = {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    textAlign: 'center',
    minHeight: '50vh',
}

const IndexPage = () => {
    useEffect(() => {
        const token = localStorage.getItem('token')
        if (!token) {
            navigate('/auth/login')
        }
    }, [])
    
    return (
        <main style={pageStyles}>
            <title>e-bpm</title>
            <h1 style={headingStyles}>Aplikasi Bidan Praktek Mandiri (BPM)</h1>
        </main>
    )
}

export default IndexPage