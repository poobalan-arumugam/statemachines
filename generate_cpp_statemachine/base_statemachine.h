#ifndef __MPA__BASE_STATE_MACHINE__HPP__
#define __MPA__BASE_STATE_MACHINE__HPP__

#include <stdexcept>

#define SM_UNUSED_ARG(arg) \
{ (void)arg; }

namespace nsii {
namespace statemachine {


///
/// \brief The Event class
///
/// Event does not need to be used by an implementation.
/// As long as the implementation provides the _typecode method.
///
class Event
{
public:
    virtual ~Event(){}

    virtual int _typecode() const = 0;
    //static int _cpp_typecode() { return 0; }
};


std::ostream& operator <<(std::ostream& os, const Event& ev)
{
    os << "Event? " << ((void*)&ev);
    return os;
}

///
/// \brief The EventSource class
///
/// EventSource must be derived from using the CRTP pattern
/// to declare its first template type.
/// (http://en.wikipedia.org/wiki/Curiously_recurring_template_pattern)
///
/// e.g. class ABC : public nsii::statemachine::EventSource<ABC, MyEvent>{};
///
template <typename EventSourceBase, typename EventBase>
class EventSource
{
public:
    virtual ~EventSource(){}

    virtual void send(EventSourceBase& source, const EventBase& ev)
    {
        SM_UNUSED_ARG(source);
        SM_UNUSED_ARG(ev);
        throw std::logic_error("Must be overriden");
    }
};

template <typename EventSourceBase, typename EventBase>
std::ostream& operator <<(std::ostream& os,
                          const EventSource<EventSourceBase, EventBase>& source)
{
    os << "EventSource? " << ((void*)&source);
    return os;
}

///
/// \brief The PortBinding class
///
/// PortBinding acts as an event source on behalf of a
/// real event source.
///
template <typename Owner, typename Target, typename EventBase>
class PortBinding
{
public:
    PortBinding()
        : _owner(),
          _target()
    {}

    void bind(Owner& owner, Target& target)
    {
        _owner = &owner;
        _target = &target;
    }

    void bind(Owner* owner, Target* target)
    {
        _owner = owner;
        _target = target;
    }

    void send(const EventBase& ev)
    {
        if(nullptr == _target)
        {
            throw std::logic_error("Target for Port not bound");
        }
        if(nullptr == _owner)
        {
            throw std::logic_error("Owner for Port not bound");
        }
        _target->send(*_owner, ev);
    }

private:
    Owner* _owner;
    Target* _target;
};

///
/// \brief The statemachine_t class
///
/// Note that statechine_t is fully templatized on the event source
/// and event base types - with the event base having a requirement
/// as in "Event" above.
///
/// One way of getting a derived implementation is to use the
/// StateProto project and the generate_cpp.py scratch file.
/// This will be improved over time.
///
template <typename EventSourceBase, typename EventBase>
class statemachine_t
{
public:
    statemachine_t(const std::string& instanceid)
        : _is_initialised(),
          _state(),
          _instanceid(instanceid)
    {}

    virtual void initialise() = 0;
    virtual void dispatch(EventSourceBase& source, const EventBase& ev) = 0;

    virtual std::string modelname() const = 0;
    virtual std::string modelnamespace() const = 0;
    virtual std::string modelguid() const = 0;
    virtual std::string modelfilename() const = 0;
    virtual std::string statemachineversion() const = 0;
    virtual std::string comment() const = 0;

    std::string instanceid() const { return _instanceid; }
    void instanceid(const std::string& value) { _instanceid = value; }

    inline int state() const { return _state; }
    inline void set_state(int value){ _state = value; }

    inline bool is_initialised() const { return _is_initialised; }
    inline void set_initialised() { _is_initialised = true; }

    virtual void unhandled_event(const EventSourceBase& source, const EventBase& ev)
    {
        _model_unhandled_event(source, ev);
    }

protected:
    virtual void _model_unhandled_event(
            const EventSourceBase& source,
            const EventBase& ev) = 0;

private:
    bool _is_initialised;
    int _state;
    std::string _instanceid;
};

} // statemachine
} // nsii

#endif // __MPA__BASE_STATE_MACHINE__HPP__
