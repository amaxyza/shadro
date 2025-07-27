import React, { useState } from "react"
import { useParams } from "react-router-dom"
import Header from "../components/Header"
import GlslEditor from "../components/GlslEditor"

const Create: React.FC = () => { 
    const { id } = useParams()
    return (
        
        <>
            <Header />
            <GlslEditor program_id={id ? parseInt(id, 10) : undefined}></GlslEditor>
        </>
    )
}

export default Create