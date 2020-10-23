import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser'; // import interface User 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 //const [users, setUsers] = useState(Array);
 const [loading, setLoading] = useState(true);

 const [users1, setUsers1] = React.useState<EntUser[]>([]);
 
 useEffect(() => {
   const getUsers = async () => {
     const res = await api.listUser({ limit: 20, offset: 0 });
     setLoading(false);
     setUsers1(res);
     console.log(res);
   };
   getUsers();
 }, [loading]);

 const deleteUsers = async (id: number) => {
  const res = await api.deleteUser({ id: id });
  setLoading(true);
};
console.log(users1)
 
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
         <TableCell align="center">ลำดับที่</TableCell>
          <TableCell align="center">ประเภทของ User</TableCell>
          <TableCell align="center">เลขประจำตัวประชาชน</TableCell>
           <TableCell align="center">รหัสผ่าน</TableCell>
           <TableCell align="center">ยืนยันรหัสผ่าน</TableCell>
           <TableCell align="center">ชื่อ</TableCell>
           <TableCell align="center">นามสกุล</TableCell>
           <TableCell align="center">เพศ</TableCell>
           <TableCell align="center">อีเมล</TableCell>
           <TableCell align="center">จังหวัด</TableCell>
           <TableCell align="center">เบอร์โทรศัพท์</TableCell>
           <TableCell align="center">วันเดือนปีเกิด</TableCell>
           <TableCell align="center">delete</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
              
         
         {users1.map((item:any) => (
           <TableRow key={item.id}>
              <TableCell align="center">{item.id}</TableCell>
            <TableCell align="center">{item.edges?.userTypeID?.userTypeName}</TableCell>
            <TableCell align="center">{item.identityCard}</TableCell>
            <TableCell align="center">{item.password}</TableCell>
            <TableCell align="center">{item.confirmPassword}</TableCell>
            <TableCell align="center">{item.firstName}</TableCell>
            <TableCell align="center">{item.lastName}</TableCell>
             <TableCell align="center">{item.edges?.genderID?.genderName}</TableCell>
             <TableCell align="center">{item.email}</TableCell>
             <TableCell align="center">{item.edges?.provinceID?.provinceName}</TableCell>
             <TableCell align="center">{item.phone}</TableCell>
             <TableCell align="center">{item.dateOfBirth}</TableCell>
           
             <TableCell align="center">
             <Button
                 onClick={() => {
                   deleteUsers(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
