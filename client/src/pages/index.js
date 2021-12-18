import React, { useEffect } from 'react';
import { navigate } from 'gatsby';

const IndexPage = () => {
    useEffect(() => {
        const token = localStorage.getItem('token');

        if (!token) {
            navigate('/auth/login');
        }
    }, []);
    
    return 
}

export default IndexPage
