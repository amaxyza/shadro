import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import './index.css'

import Home from './pages/Home'
import Login from './pages/Login'
import Signup from './pages/Signup'
import Create from './pages/Create'
import ProfilePage from './pages/ProfilePage'

createRoot(document.getElementById('root')!).render(
    <BrowserRouter>
        <Routes>
            <Route path='/' element={<Home />} />
            <Route path='/login' element={<Login />} />
            <Route path='/signup' element={<Signup />} />
            <Route path='/create' element={<Create />} />
            <Route path='/about' element={<Home />} />
            <Route path="/profiles/:id" element={<ProfilePage />} />

        </Routes>
    </BrowserRouter>
)
