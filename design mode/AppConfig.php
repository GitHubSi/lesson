<?php

//工厂模式还是单例模式

class AppConfig
{
    private static $instance;
    private $commsManage;

    private function __construct()
    {

    }

    private function init()
    {
        switch (Setting::$COMMSTYPE) {
            case "Mega" :
                $this->commsManage = new MegaCommsManager();
                break;
            default:
                $this->commsManage = new BloggaCommsManager();
        }
    }

    public static function getInstance()
    {
        if (!empty(self::$instance)) {
            self::$instance = new self();
        }
        return self::$instance;
    }
}