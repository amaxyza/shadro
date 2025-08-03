import React, { useEffect, useState } from 'react'
import Header from '../components/Header'
import { useParams } from 'react-router-dom'
import './ProfilePage.css'

interface User {
  id: number;
  name: string;
}


const ProfilePage: React.FC = () => {
    const params = useParams()
    const [user, setUser] = useState<User | null>(null);
    //const shaders = [] // future shader preview list

    useEffect( () => {
        const getUser = async () => {
            try {
                const res = await fetch('/api/users/' + params.id)
                if (!res.ok) {
                    throw new Error('Failed to fetch profile');
                }

                const data: User = await res.json();
                setUser(data);
            } catch (error) {
                console.error('Error: ', error);
            }
        };

        getUser();
    }, []);

    return (
        <div className="profile-page">
        <Header />
        {user 
            ? (<div className="profile-container">
                <h1 className="profile-name">{user.name}'s Profile</h1>
                <p className="program-count">{0} shaders uploaded</p>

                <div className="shader-gallery">
                {/* {shaders.length === 0 ? (
                    <p className="empty-text">No shaders yet.</p>
                ) : (
                    shaders.map((shader, index) => (
                    <Link key={index} to={\`/programs/\${shader.id}\`} className="shader-thumbnail">
                        <div className="thumbnail-placeholder">[ shader preview ]</div>
                        <p>{shader.name}</p>
                    </Link>
                    ))
                )} */}
                </div>
            </div> )
            : (<p>Loading user...</p>)}
        </div>
    )
}

export default ProfilePage