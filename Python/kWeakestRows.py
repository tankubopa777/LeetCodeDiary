def kWeakestRows(mat, k):
    # นับจำนวน 1 ในแต่ละแถวและรักษาดัชนีของแถว
    counts = [(sum(row), idx) for idx, row in enumerate(mat)]
    print(counts)
    # เรียงลำดับตามจำนวน 1 และหากจำนวนเท่ากันจะเรียงตามดัชนีแถว
    sorted_counts = sorted(counts)
    print(sorted_counts)
    
    # เลือก k ดัชนีแถวที่มีจำนวน 1 น้อยที่สุด
    weakest_rows = [idx for _, idx in sorted_counts[:k]]
    print(weakest_rows)
    
    return weakest_rows

# ทดสอบฟังก์ชัน
mat = [[1,1,0,0,0],
       [1,1,0,0,0],
       [1,1,0,1,0],
       [1,0,0,0,0]]

k = 5
print(kWeakestRows(mat, k))