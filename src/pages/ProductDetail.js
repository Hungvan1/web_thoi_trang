import React, { useEffect, useState, useContext } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { CartContext } from '../context/CartContext';

const ProductDetail = () => {
    const { id } = useParams();
    const [prod, setProd] = useState(null);
    const [qty, setQty] = useState(1);
    const { addToCart } = useContext(CartContext);
    const navigate = useNavigate();

    useEffect(() => {
        // Replace with your new backend API call to fetch product details
        setProd({ 
            id: id, 
            title: `Product ${id}`, 
            price: 250000, 
            description: 'This is a sample product description.', 
            image: 'https://via.placeholder.com/150' 
        }); // Simulate a product for now
    }, [id]);

    if (!prod) return <div>Loading...</div>;

    return (
        <div className="product-detail container">
            <div className="pd-left">
                <img src={prod.image} alt={prod.title} />
            </div>
            <div className="pd-right">
                <h2>{prod.title}</h2>
                <p className="price">
                    {new Intl.NumberFormat('vi-VN', { style: 'currency', currency: 'VND' }).format(prod.price)}
                </p>
                <p>{prod.description}</p>

                <div>
                    <label>Size: </label>
                    <select>
                        <option>S</option><option>M</option><option>L</option><option>XL</option>
                    </select>
                </div>

                <div className="pd-actions">
                    <input type="number" min="1" value={qty} onChange={(e) => setQty(Number(e.target.value))} />
                    <button onClick={() => { addToCart(prod, qty, { size: 'M' }); alert('Added to cart'); }}>Thêm vào giỏ</button>
                    <button onClick={() => { addToCart(prod, qty, { size: 'M' }); navigate('/cart'); }}>Mua ngay</button>
                </div>
            </div>
        </div>
    );
};

export default ProductDetail;
