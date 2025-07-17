import React, { useState } from 'react'
import { Link, Navigate, useNavigate } from 'react-router-dom' 

const Login: React.FC = () => {
  const [name, setName] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const navigate = useNavigate()

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    const res = await fetch('/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name, password }),
    })
    if (!res.ok) {
      setError('Login failed.')
    } else {
      // Redirect or handle success
      const data = await res.json()
      if (data.success) {
        console.log('Logged in!')
        navigate('/profiles/' + data.id)
      } else {
        console.error('failed to convert response to json')
      }
      
    }
  }

  return (
    <div style={styles.container}>
      <h2>Login</h2>
      <form onSubmit={handleSubmit} style={styles.form}>
        <input
          type="text"
          placeholder="Username"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
          style={styles.input}
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
          style={styles.input}
        />
        <button type="submit" style={styles.button}>Login</button>
        {error && <p style={styles.error}>{error}</p>}
      </form>
      <p>Don't have an account?</p>
      <Link to='/signup'>Create an Account</Link>
    </div>
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
  form: { display: 'flex', flexDirection: 'column', gap: 10, width: 300 },
  input: { padding: 10, fontSize: 16 },
  button: { padding: 10, fontSize: 16, cursor: 'pointer' },
  error: { color: 'red', marginTop: 10 },
}

export default Login
