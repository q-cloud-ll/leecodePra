#进店却未进行过交易的顾客
select customer_id, count(customer_id) as count_no_trans
from Visits v left join Transactions t
on v.visit_id = t.visit_id
where t.amount is NULL
group by customer_id;

#部门工资前三高的员工
select e.Department,
       e.Employee,
       e.Salary
from (
     select
         Department.Name as 'Department',
         Employee.Name as 'Employee',
         Employee.Salary,
         dense_rank() over(partition by Employee.DepartmentId order by Employee.Salary desc) as 'Ranking'
     from Employee
               join Department
                         on Employee.DepartmentId = Department.Id
 ) as e
where e.Ranking <= 3


select max(Salary) SecondHighestSalary
from Employee
where Salary < (select max(Salary) from Employee)