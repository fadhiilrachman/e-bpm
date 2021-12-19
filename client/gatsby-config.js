module.exports = {
  siteMetadata: {
    siteUrl: `http://localhost:8000`,
    title: `e-bpm`,
    description: `Aplikasi Bidan Praktek Mandiri`,
    author: `Fadhiil Rachman`,
  },
  plugins: [
    {
      resolve: `gatsby-plugin-manifest`,
      options: {
        icon: `src/images/icon.png`,
      },
    },
  ],
};
