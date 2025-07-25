import React from 'react'
import { Link } from 'react-router-dom'
import Header from '../components/Header'

const Home: React.FC = () => {
  return (
    <>
      <Header />
      <div style={styles.container}>
        <h1>Welcome to Shadro</h1>
        <p style={styles.subtext}>A place to share and explore GLSL shaders.</p>
        {/* <div style={styles.links}>
          <Link to="/login" style={styles.linkButton}>Login</Link>
          <Link to="/signup" style={styles.linkButton}>Sign Up</Link>
        </div> */}
      </div>
    </>
  )
}

const styles: { [key: string]: React.CSSProperties } = {
    container: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    minHeight: '100vh',
    width: '100%',
    padding: 40,
    boxSizing: 'border-box',
    }, 
  subtext: {
    marginTop: 10,
    fontSize: 18,
    color: '#666',
  },
  links: {
    display: 'flex',
    gap: 20,
    marginTop: 30,
  },
  linkButton: {
    padding: '10px 20px',
    backgroundColor: '#333',
    color: '#fff',
    textDecoration: 'none',
    borderRadius: 5,
    fontSize: 16,
  },
}

export default Home
