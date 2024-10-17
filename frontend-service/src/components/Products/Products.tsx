import { useEffect, useState } from "react";
import ProductCard from "./ProductCard";
import Pagination from "./Pagination";

export type Product = {
  _id?: string
  name: string
  brand: string
  content: string
  added_by: string
}

type Props = {
  checkAdded:boolean
}

const Products = ({checkAdded}:Props) => {
  const [currentPage,setCurrentPage]=useState<number>(1)
  const [totalPage,setTotalPage]=useState<number>(0)
  const [products, setProducts] = useState<Product[]>([{
    _id: "",
    name: "",
    brand: "",
    content: "",
    added_by: "",
  }])

    const getProducts = async(page:number) => {
        const response = await fetch(`http://localhost:8083/api/v1/products?page=${page}`)
      const res = await response.json()
      console.log(res);
      setProducts(res.data) 
      setCurrentPage(res.page)
      setTotalPage(res.totalPage)
    }
    useEffect(() => {
        getProducts(currentPage)
    },[currentPage,checkAdded])
  return (
    <>
    <div className="flex flex-wrap gap-2 justify-between p-2 ">
      {
        products.map((product:Product) =>
          {           
            return <ProductCard key={product._id} props={product} /> 
          }
        )}
    </div>

     <Pagination totalPages={totalPage} currentPage={currentPage} onPageChange={setCurrentPage} />
        </>
  )
}

export default Products