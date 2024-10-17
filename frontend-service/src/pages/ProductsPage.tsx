import React, { useEffect, useState } from 'react'
import Products from '../components/Products/Products';
import AddProduct from '../components/Products/AddProduct';

const ProductsPage: React.FC = () => {
  const [addProduct, setAddProduct] = useState<boolean>(false)
  const [added,setAdded] = useState<boolean>(false)
  
  const handleAddProduct = () => {
    setAddProduct(true)
  }
  
  const handleAdded = () => {
    setAdded(!added)
  }
    
  return (
    <div className='relative'>
      <button onClick={handleAddProduct}>Add Product </button>
      {addProduct && <AddProduct handleAdded={handleAdded} />}
        
     
      <Products checkAdded={added} />
    </div>
  )
}

export default ProductsPage