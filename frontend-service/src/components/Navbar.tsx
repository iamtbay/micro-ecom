import { useEffect, useState } from "react";

export const checkCookie = async () => {
  const response = await fetch("http://localhost:8080/api/v1", {
    method: "POST",
    body: JSON.stringify({
      "action": "auth",
      "sub_action":"check"
    }),
  credentials:"include"
  })
  const res = await response.json()
  console.log(res);
  
if (!response.ok) {
  if (response.status === 401) {
    throw new Error("Unauthorized access")
  } else {
    throw new Error("Something went wrong")
  }
}
return true
}

function Navbar() {
  const [isLoggedIn,setIsLoggedIn]=useState<boolean>(false)
 
  
  async function checker() {
    const res = await checkCookie()
    if (res) {
      setIsLoggedIn(true)
    }
    
  } 
  
  const handleLogout = async () => {
    const response = await fetch("http://localhost:8080/api/v1", {
      method: "POST",
      body: JSON.stringify({
        "action": "auth",
        "sub_action":"logout",
      }),
      credentials:'include'
    })
    const result = await response.json()
    console.log(result);
    
  }
  useEffect(() => {
    checker()
  }, [])
  
  return (
      <div className='flex flex-row bg-blue-200 justify-center p-4'>
          <ul className='flex flex-row gap-4'>
              <li><a className="hover:text-blue-700 transition duration-100" href="/">Mainpage</a></li>
              <li><a className="hover:text-blue-700 transition duration-100" href="/products">Products</a></li>
              {!isLoggedIn && <li><a className="hover:text-blue-700 transition duration-100" href="/auth">Authentication</a></li>}
        {
          isLoggedIn &&
          <li className="hover:text-blue-700 transition duration-100 cursor-pointer" onClick={handleLogout}>Logout</li>}
          </ul>
    </div>
  )
}

export default Navbar