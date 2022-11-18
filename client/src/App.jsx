import { useEffect, useState } from 'react'

function App() {
    
  const [name, setName] = useState('');
  const [certs, setCerts] = useState([])

  async function loadCerts(){
    const response = await fetch(/*import.meta.env.VITE_BE_URI + */'/certs')
    const data = await response.json()
    setCerts(data.certs)
  }

  useEffect(() => {
    loadCerts()
  }, [])

  const handleSubmit = async (e) => {
    e.preventDefault()
    const response = await fetch(/*import.meta.env.VITE_BE_URI + */'/certs', {
      method: 'POST',
      body: JSON.stringify({name}),
      headers: {
        "Content-Type" : "application/json"
      }
    })
    const data = await response.json()
    loadCerts()
  }

  return (
    <div>

    <form onSubmit={handleSubmit}>
      <input 
        type="name" 
        placeholder="Certification name" 
        onChange={e => setName(e.target.value)} 
      />

      <button>Save</button>
    </form>

    <ul>
      {certs.map(cert => (
        <li key={cert._id}>{cert.name}</li>
      ))}
    </ul>

    </div>
  )
}

export default App
