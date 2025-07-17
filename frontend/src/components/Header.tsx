import React, { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'

interface User {
  id: number;
  username: string;
}

const Header: React.FC = () => {
  const [username, setUsername] = useState<User | null | undefined>(undefined)

  useEffect( () => {
    const getName = async () => {
      try {
        const res = await fetch('/api/me', {credentials: 'include'})
        if (res.ok) {
          const data: User = await res.json();
          setUsername(data);
        }
        else {
          console.log("failed to get login")
          setUsername(null)
        }

      } catch (error) {
        console.log('Error: ', error)
      }
    };

    getName();
  }, [])

  return (
    <header style={styles.header}>
      <div style={styles.container}>
        <div style={styles.left}>
          <Link style={styles.link} to="/about">About</Link>
          <Link style={styles.link} to ="/create">Create</Link>
        </div>
        
        <div style={styles.right}>
          {username === undefined ? (
            null
          ) : username ? (
            <Link style={styles.link} to={"/profiles/" + username.id}>Hello, {username.username}!</Link>
          ) : (
            <>
              <Link style={styles.link} to="/signup">Signup</Link>
              <Link style={styles.link} to="/login">Login</Link>
            </>
          )}
        </div>
      </div>
    </header>
  )
}

const styles: { [key: string]: React.CSSProperties } = {
  header: {
    backgroundColor: '#393939ff',
    padding: '10px 20px',
    borderBottom: '1px solid #ccc',
  },
  container: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
  },
  left: {
    display: 'flex',
    gap: '20px',
  },
  right: {
    display: 'flex',
    gap: '20px',
  },
  link: {
    textDecoration: 'none',
    color: '#fff',
    fontWeight: 'bold',
  },
}

export default Header
