import React, { Dispatch, useState } from 'react'
import InputComponent from '../InputComponent';
import { Product } from './Products';

type Props = {
    handleAdded:()=>void
}

const AddProduct = ({handleAdded}:Props) => {

   
    const [newProduct,setNewProduct] =useState<Product>({
        name:'',
        brand:"",
        content:'',
        added_by:""
      })
    const handleSubmit = async (e:React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const response = await fetch("http://localhost:8083/api/v1/products/add", {
          method: "POST",
          body: JSON.stringify(newProduct),
          credentials:'include'
        })
        const result = await response.json()
        handleAdded()
    }
    
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setNewProduct((val) => (
               { ...val,
                [e.target.name]:e.target.value}
        ))
    }

  return (
      <div className='absolute bg-gray-500 w-full p-2'>
          
    <form onSubmit={handleSubmit}>

              <InputComponent name="name" placeholder="x jewelry" labelText='Product Name' type='text' value={newProduct.name} onChange={handleChange} />

              <InputComponent name="brand" placeholder="x brand" labelText='Product Brand' type='text' value={newProduct.brand} onChange={handleChange} />

              <InputComponent name="content" placeholder="example of the product content with details" labelText='Product Content' type='text' value={newProduct.content} onChange={handleChange} />

              <input className='bg-green-200 p-2 rounded ' type="submit" value="Add Product" />
      </form> 

    </div>
  )
}

export default AddProduct