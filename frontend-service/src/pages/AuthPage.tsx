import React, { useEffect, useState } from 'react'
import InputComponent from '../components/InputComponent'
import { checkCookie } from '../components/Navbar'
import { Navigate, redirect, useNavigate } from 'react-router-dom'

const AuthPage: React.FC = () => {
  const navigate = useNavigate()
  const [authMethod, setAuthMethod] = useState<string>("login")
  const [userInfo, setUserInfo] = useState({
    name: "",
    surname: "",
    email: "",
    password: "",
  })
  // fix the height problem thats when click login it's 
  // dropping to 2 rows and the button and information
  // text changing their place
  const changeMethod = () => {
    if (authMethod === "login") {
      setAuthMethod("register")
    } else if (authMethod === "register") {
      setAuthMethod("login")
    }
  }

  const changeUserInfo = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUserInfo((values) =>
    ({
      ...values,
      [e.target.name]: e.target.value
      
    }))
  }

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const method = authMethod == "login" ? "login" : "signup"
    var response = await fetch(`http://localhost:8080/api/v1`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        "action": "auth",
        "sub_action":method,
        "data": userInfo
      }),
      credentials: "include"
    })

    var result = await response.json()
    if (!response.ok) {
      console.log(result);
      
      throw new Error(result.error)
    }
    resetUserInfo()
    console.log(result);

  }

  const resetUserInfo = () => {
    setUserInfo(() => ({
      name: "",
      surname: "",
      email: "",
      password: ""
    }))
  }
  async function checker() {
    const auth = await checkCookie()
    console.log(auth);
    
    if (auth) {
        navigate("/")
    }
  }

  useEffect(() => {
    checker()
  },[])

  
  return (
    <form onSubmit={handleSubmit} className='w-full min-h-[360px] p-2 flex flex-col justify-between items-center gap-1 bg-red-200'>
      <input type="button" value="reset" onClick={resetUserInfo} />
      <div className='flex flex-col w-full items-center '>
      {
        authMethod == "register" &&
        <>
            <InputComponent name="name" value={userInfo.name} placeholder='john' onChange={changeUserInfo} />
         <InputComponent name="surname" value={userInfo.surname}  placeholder='doe' onChange={changeUserInfo} />
        </>
      }
      <InputComponent name="email" value={userInfo.email}  placeholder='johndoe@gmail.com' onChange={changeUserInfo} />
      <InputComponent name="password" value={userInfo.password}  placeholder='password' type="password" onChange={changeUserInfo} />
      
      </div>
      <div className='flex flex-col items-end items-center self-center'>

      {
        authMethod === "login" ?
        <p>Don't have an account? <span className='text-blue-500 cursor-pointer' onClick={changeMethod}>Sign up here</span></p>
        :
        <p>Already have an account? <span className='text-blue-500 cursor-pointer' onClick={changeMethod}>Login here</span></p>
      }

      <input type='submit' className='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded' value={ 
        authMethod === "login" ?
        "Login" :
        "Register"
      } />
       
       
      </div>
    </form>
  )
}

export default AuthPage