import React, { useEffect, useState } from 'react';
import ProductCard from '../components/ProductCard';
import { useLocation } from 'react-router-dom';

const Products = () => {
    const [products, setProducts] = useState([]);
    const location = useLocation();
    const searchParams = new URLSearchParams(location.search);
    const searchQuery = searchParams.get('search') || '';

    const [gender, setGender] = useState('all');
    const [size, setSize] = useState('all');

    useEffect(() => {
        // Replace with your new backend API call to fetch products
        setProducts([]); // Initialize with an empty array
    }, []);

    const filtered = products.filter(p => {
        // Filter by search query from Navbar (exact match)
        const matchesSearchQuery = searchQuery ? p.title.toLowerCase() === searchQuery.toLowerCase() : true;

        // Filter by gender based on title includes
        const matchesGender = gender === 'all'
            || (gender === 'men' && p.title.toLowerCase().includes("nam"))
            || (gender === 'women' && p.title.toLowerCase().includes("nữ"));

        // giả lập size (FakeStoreAPI không có size, ta random assign size pool)
        const fakeSizes = ['S', 'M', 'L', 'XL'];
        const assignedSize = fakeSizes[p.id % fakeSizes.length];
        const matchesSize = size === 'all' || assignedSize === size;

        return matchesSearchQuery && matchesGender && matchesSize;
    });

    return (
        <div className="container">
            <h2>Sản phẩm</h2>
            <div className="filter-bar">
                {/* <input placeholder="Tìm kiếm..." value={q} onChange={e => setQ(e.target.value)} /> */}
                <select value={gender} onChange={e => setGender(e.target.value)}>
                    <option value="all">Tất cả</option>
                    <option value="men">Nam</option>
                    <option value="women">Nữ</option>
                </select>
                <select value={size} onChange={e => setSize(e.target.value)}>
                    <option value="all">Mọi size</option>
                    <option value="S">S</option>
                    <option value="M">M</option>
                    <option value="L">L</option>
                    <option value="XL">XL</option>
                </select>
            </div>

            <div className="product-grid">
                {filtered.map(p => <ProductCard key={p.id} product={p} />)}
            </div>
        </div>
    );
};

export default Products;
