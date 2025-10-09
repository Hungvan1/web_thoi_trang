import React from 'react';
import { Link } from 'react-router-dom';

const ProductCard = ({ product }) => {
    const discount = Math.floor(Math.random() * 45) - 5;
    return (
        <div className="product-card">
            <div className="badge">{discount}%</div>
            <Link to={`/product/${product.id}`}>
                <img src={product.image} alt={product.title} />
            </Link>
            <div className="prod-info">
                <h4>{product.title}</h4>
                <p className="price">
                    {new Intl.NumberFormat('vi-VN', { style: 'currency', currency: 'VND' }).format(product.price)}
                </p>
            </div>
        </div>
    );
};

export default ProductCard;
