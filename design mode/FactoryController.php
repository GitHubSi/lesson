<?php

//  组合模式中的一切类都共享同一个父类。客户端也没有必要区分一个对象是组合对象还是局部对象。
//但也意味这局部类也必须实现父类中的方法，但有些方法却是没有必要的。

abstract class Unit
{
    abstract function addUnit(Unit $unit);

    abstract function removeUnit(Unit $unit);

    //修改基类的方法，在基类中抛出异常，省去了局部类中变量的声明，去除冗余的类方法 -002
    //function removeUnit(Unit $unit)
    //{
    //    throw new UnitException(get_class($this) . "is a leaf");
    //}

    abstract function bombardStrength();
}

class Army extends Unit
{
    private $units = array();

    function addUnit(Unit $unit)
    {
        if (in_array($unit, $this->units, true)) {
            return;
        }
        $this->units[] = $unit;
    }

    function removeUnit(Unit $unit)
    {
        $this->units = array_udiff($this->units, array($unit));
    }

    function bombardStrength()
    {
        $ret = 0;
        foreach ($this->units as $unit) {
            $ret += $unit->bombardStrength();
        }
        return $ret;
    }
}

class UnitException extends Exception
{

}

//因为局部类中不是允许调用add和remove方法的，我们也可以在设计之前，在基类中直接抛出异常 -001
class Archer extends Unit
{
    function addUnit(Unit $unit)
    {
        throw new UnitException(get_class($this) . "is a leaf");
    }

    function removeUnit(Unit $unit)
    {
        throw new UnitException(get_class($this) . "is a leaf");
    }

    function bombardStrength()
    {
        return 4;
    }
}

//修改一下继承结构,这样也省去了组合类中冗余的方法
abstract class CompositeUnit extends Unit
{
    private $units = array();

    function getComposite()
    {
        return $this;
    }

    function addUnit(Unit $unit)
    {
        if (in_array($unit, $this->units, true)) {
            return;
        }
        $this->units[] = $unit;
    }

    function removeUnit(Unit $unit)
    {
        $this->units = array_udiff($this->units, array($unit));
    }
}