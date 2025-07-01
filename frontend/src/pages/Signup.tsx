import React, { useState } from 'react'
import { Link } from 'react-router-dom'

const Signup: React.FC = () => {
  const [name, setName] = useState('')
  const [password, setPassword] = useState('')
  const [message, setMessage] = useState('')

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setMessage('')
    const res = await fetch('/api/signup', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name, password }),
    })
    if (!res.ok) {
      setMessage('Signup failed.')
    } else {
      setMessage('Account created!')
    }
  }

  return (
    <div style={styles.container}>
      <h2>Sign Up</h2>
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
        <button type="submit" style={styles.button}>Sign Up</button>
        {message && <p>{message}</p>}
      </form>
      <p>Have an account?</p>
      <Link to='/login'>Login to your account</Link>
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
}

export default Signup
