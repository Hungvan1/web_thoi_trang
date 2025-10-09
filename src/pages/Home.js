import React, { useEffect, useState } from 'react';
import ProductCard from '../components/ProductCard';
import { Link } from 'react-router-dom';

const Home = () => {
    const [featured, setFeatured] = useState([]);

    useEffect(() => {
        // Replace with your new backend API call to fetch featured products
        setFeatured([]); // Initialize with an empty array
    }, []);

    return (
        <div>
            <section className="hero">
                <h1>THỜI TRANG HOT NHẤT</h1>
                <p>Website demo bán quần áo</p>
            </section>

            <section className="featured-grid">
                {featured.map(p => <ProductCard key={p.id} product={p} />)}
            </section>

            <div style={{ textAlign: 'center', margin: '30px' }}>
                <Link to="/products" className="btn">Xem tất cả sản phẩm</Link>
            </div>
        </div>
    );
};

export default Home;
